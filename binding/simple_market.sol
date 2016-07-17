contract EventfulMarket {
    event ItemUpdate( uint id );
    event Trade( uint sell_how_much, address indexed sell_which_token,
                 uint buy_how_much, address indexed buy_which_token );
}

contract FallbackFailer {
  function () {
    throw;
  }
}

contract Assertive {
  function assert(bool assertion) internal {
    if (!assertion) throw;
  }
}

contract ERC20 {
    function totalSupply() constant returns (uint);
    function balanceOf(address who) constant returns (uint);
    function allowance(address owner, address spender) constant returns (uint);

    function transfer(address to, uint value) returns (bool ok);
    function transferFrom(address from, address to, uint value) returns (bool ok);
    function approve(address spender, uint value) returns (bool ok);

    event Transfer(address indexed from, address indexed to, uint value);
    event Approval(address indexed owner, address indexed spender, uint value);
}

contract SimpleMarket is EventfulMarket, FallbackFailer, Assertive {
    struct OfferInfo {
        uint sell_how_much;
        ERC20 sell_which_token;
        uint buy_how_much;
        ERC20 buy_which_token;
        address owner;
        bool active;
    }
    uint public last_offer_id;
    mapping( uint => OfferInfo ) public offers;

    function next_id() internal returns (uint) {
        last_offer_id++; return last_offer_id;
    }
    function offer( uint sell_how_much, ERC20 sell_which_token
                  , uint buy_how_much,  ERC20 buy_which_token )
        returns (uint id)
    {
        assert(sell_how_much > 0);
        assert(sell_which_token != ERC20(0x0));
        assert(buy_how_much > 0);
        assert(buy_which_token != ERC20(0x0));

        var seller_paid = sell_which_token.transferFrom( msg.sender, this, sell_how_much );
        assert(seller_paid);
        OfferInfo memory info;
        info.sell_how_much = sell_how_much;
        info.sell_which_token = sell_which_token;
        info.buy_how_much = buy_how_much;
        info.buy_which_token = buy_which_token;
        info.owner = msg.sender;
        info.active = true;
        id = next_id();
        offers[id] = info;
        ItemUpdate(id);
        return id;
    }
    function trade( address seller, uint sell_how_much, ERC20 sell_which_token, address buyer, uint buy_how_much, ERC20 buy_which_token ) internal
    {
        var seller_paid_out = buy_which_token.transferFrom( buyer, seller, buy_how_much );
        assert(seller_paid_out);
        var buyer_paid_out = sell_which_token.transfer( buyer, sell_how_much );
        assert(buyer_paid_out);
        Trade( sell_how_much, sell_which_token, buy_how_much, buy_which_token );
    }
    function buy( uint id ) returns ( bool _success )
    {
        var offer = offers[id];
        assert(offer.active);

        trade( offer.owner, offer.sell_how_much, offer.sell_which_token,
               msg.sender, offer.buy_how_much, offer.buy_which_token );

        delete offers[id];
        ItemUpdate(id);
        return true;
    }
    function buyPartial( uint id, uint quantity ) returns ( bool _success )
    {
        var offer = offers[id];
        assert(offer.active);

        if ( offers[id].sell_how_much < quantity ) {
            return false;
        } else if ( offers[id].sell_how_much == quantity ) {
            trade( offer.owner, offer.sell_how_much, offer.sell_which_token,
                   msg.sender, offer.buy_how_much, offer.buy_which_token );
            delete offers[id];
            ItemUpdate(id);
            return true;
        } else {
            uint buy_quantity = quantity * offers[id].buy_how_much / offers[id].sell_how_much;
            if ( buy_quantity > 0 ) {
                trade( offer.owner, quantity, offer.sell_which_token,
                       msg.sender, buy_quantity, offer.buy_which_token );

                offer.sell_how_much -= quantity;
                offer.buy_how_much -= buy_quantity;
                ItemUpdate(id);
                return true;
            }
        }
    }
    function cancel( uint id ) returns ( bool _success )
    {
        var offer = offers[id];
        assert(offer.active);
        assert(msg.sender == offer.owner);

        var seller_refunded = offer.sell_which_token.transfer( msg.sender, offer.sell_how_much );
        assert(seller_refunded);
        delete offers[id];
        ItemUpdate(id);
        return true;
    }
    function getOffer( uint id ) constant
        returns (uint, ERC20, uint, ERC20) {
      var offer = offers[id];
      return (offer.sell_how_much, offer.sell_which_token,
              offer.buy_how_much, offer.buy_which_token);
    }
}
