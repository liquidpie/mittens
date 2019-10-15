//Copyright 2019 Expedia, Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package flags

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Profile struct {
	CPU    string
	Memory string
}

func (p *Profile) String() string {
	return fmt.Sprintf("%+v", *p)
}

func (p *Profile) InitFlags(cmd *cobra.Command) {

	cmd.Flags().StringVar(
		&p.CPU,
		"profile-cpu",
		"",
		"Name of the file where to write CPU profile data",
	)
	cmd.Flags().StringVar(
		&p.Memory,
		"profile-memory",
		"",
		"Name of the file where to write memory profile data",
	)
}