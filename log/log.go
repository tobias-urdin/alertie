// Alertie
// Copyright (c) 2018 Tobias Urdin
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package log

import (
	"fmt"
	"os"
	"github.com/inconshreveable/log15"
)

var logger log15.Logger

func Init(name string) {
	logger = log15.New()

	//logger.SetHandler(log15.DiscardHandler())
	logger.SetHandler(log15.StreamHandler(os.Stderr, log15.LogfmtFormat()))
}

func Debug(format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf(format, v...))
}

func Info(format string, v ...interface{}) {
	logger.Info(fmt.Sprintf(format, v...))
}

func Warn(format string, v ...interface{}) {
	logger.Warn(fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	logger.Error(fmt.Sprintf(format, v...))
}

func Critical(format string, v ...interface{}) {
	logger.Crit(fmt.Sprintf(format, v...))
}

func Fatal(format string, v ...interface{}) {
	logger.Crit(fmt.Sprintf(format, v))
	os.Exit(1)
}
