// SPDX-FileCopyrightText: 2023-2025 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package tty_test

import (
	"bytes"
	"fmt"

	"cunicu.li/cunicu/pkg/tty"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Context("ANSIStripper", func() {
	expString := "Some color"

	It("can strip", func() {
		b := bytes.NewBuffer(nil)
		s := tty.NewANSIStripper(b)

		n, err := fmt.Fprintf(s, "Some %s", tty.Mods("color", tty.FgBlue))
		Expect(err).To(Succeed())
		Expect(n).To(BeNumerically("==", len(expString)))
		Expect(b.String()).To(Equal(expString))
	})
})
