/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"io"

	"github.com/spf13/cobra"
	"k8s.io/kops/cmd/kops/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	promoteLong = templates.LongDesc(i18n.T(`
	Promote a resource.`))

	promoteExample = templates.Examples(i18n.T(`
	# Promote the newest ca keypair to be the primary.
	kops promote keypair ca
`))

	promoteShort = i18n.T(`Promote a resource.`)
)

func NewCmdPromote(f *util.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "promote",
		Short:   promoteShort,
		Long:    promoteLong,
		Example: promoteExample,
	}

	// create subcommands
	cmd.AddCommand(NewCmdPromoteKeypair(f, out))

	return cmd
}
