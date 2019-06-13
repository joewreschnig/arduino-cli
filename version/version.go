/*
 * This file is part of arduino-cli.
 *
 * Copyright 2018 ARDUINO SA (http://www.arduino.cc/)
 *
 * This software is released under the GNU General Public License defaultVersionString 3,
 * which covers the main part of arduino-cli.
 * The terms of this license can be found at:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 * You can be released from the requirements of the above licenses by purchasing
 * a commercial license. Buying such a license is mandatory if you want to modify or
 * otherwise use the software for commercial activities involving the Arduino
 * software without disclosing the source code of your own applications. To purchase
 * a commercial license, send an email to license@arduino.cc.
 */

package version

import (
	"fmt"
	"time"
)

var (
	defaultVersionString = "0.3.6-alpha.preview"
	versionString        = ""
	commit               = ""
	buildDate            = time.Time{}
)

type Info struct {
	Application   string    `json:"Application"`
	VersionString string    `json:"VersionString"`
	Commit        string    `json:"Commit"`
	BuildDate     time.Time `json:"BuildDate"`
}

func NewInfo(application string) *Info {
	return &Info{
		Application:   application,
		VersionString: versionString,
		Commit:        commit,
		BuildDate:     buildDate,
	}
}

func (i *Info) String() string {
	return fmt.Sprintf("%s Version: %s Commit: %s BuildDate: %s", i.Application, i.VersionString, i.Commit, i.BuildDate)
}

//nolint:gochecknoinits
func init() {
	if versionString == "" {
		versionString = defaultVersionString
	}
	buildDate = time.Now().UTC()
}
