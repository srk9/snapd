// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2014-2015 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"launchpad.net/snappy/logger"
	"launchpad.net/snappy/priv"
	"launchpad.net/snappy/snappy"
)

type cmdBooted struct {
}

func init() {
	_, err := parser.AddCommand("booted",
		"internal",
		"internal",
		&cmdBooted{})
	if err != nil {
		logger.Panicf("Unable to booted: %v", err)
	}
}

func (x *cmdBooted) Execute(args []string) (err error) {
	privMutex := priv.New()
	if err := privMutex.TryLock(); err != nil {
		return err
	}
	defer privMutex.Unlock()

	parts, err := snappy.ActiveSnapsByType(snappy.SnapTypeCore)
	if err != nil {
		return err
	}

	return parts[0].(*snappy.SystemImagePart).MarkBootSuccessful()
}
