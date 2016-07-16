// Copyright (c) 2015 Max Wolter
//
// This file is part of M3 - Maker Market Maker.
//
// M3 is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// M3 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with M3.  If not, see <http://www.gnu.org/licenses/>.

package business

type fakeLogger struct {
}

func (fl *fakeLogger) Criticalf(format string, v ...interface{}) {
}

func (fl *fakeLogger) Errorf(format string, v ...interface{}) {
}

func (fl *fakeLogger) Warningf(format string, v ...interface{}) {
}

func (fl *fakeLogger) Noticef(format string, v ...interface{}) {
}

func (fl *fakeLogger) Infof(format string, v ...interface{}) {
}

func (fl *fakeLogger) Debugf(format string, v ...interface{}) {
}
