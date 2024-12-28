// SiYuan - Refactor your thinking
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/88250/gulu"
	"github.com/siyuan-note/logging"
	"golang.org/x/sys/windows"
)

// 支持一键加入 Microsoft Defender 排除项 https://github.com/siyuan-note/siyuan/issues/13650

func isUsingMicrosoftDefender() bool {
	if !gulu.OS.IsWindows() {
		return false
	}

	cmd := exec.Command("powershell", "-Command", "Get-MpPreference")
	gulu.CmdAttr(cmd)
	_, err := cmd.CombinedOutput()
	if nil != err {
		return false
	}
	return true
}

func addExclusionToWindowsDefender(exclusionPath string) {
	if !gulu.OS.IsWindows() {
		return
	}

	if isAdmin() {
		cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Add-MpPreference -ExclusionPath %s", exclusionPath))
		gulu.CmdAttr(cmd)
		output, err := cmd.CombinedOutput()
		if nil != err {
			logging.LogErrorf("add Windows Defender exclusion path [%s] failed: %s", exclusionPath, string(output))
			return
		}
	} else {
		cwd, _ := os.Getwd()
		args := strings.Join([]string{"-Command", "Add-MpPreference", "-ExclusionPath", exclusionPath}, " ")
		verbPtr, _ := syscall.UTF16PtrFromString("runas")
		exePtr, _ := syscall.UTF16PtrFromString("powershell")
		cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
		argPtr, _ := syscall.UTF16PtrFromString(args)
		var showCmd int32 = 1 //SW_NORMAL
		err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
		if err != nil {
			logging.LogErrorf("runas PowerShell failed: %s", err)
		}
	}
}

func isAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}
