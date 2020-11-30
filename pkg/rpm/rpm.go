// Package rpm は、rpm パッケージ名から 名称, バージョン などを抽出します.
package rpm

import (
	"fmt"
	"strings"
)

type (
	rpmField         string
	rpmOutputPattern string
)

// Fields
const (
	RpmName    rpmField = "name"
	RpmVersion rpmField = "version"
	RpmRel     rpmField = "rel"
	RpmArch    rpmField = "arch"
	RpmEpoch   rpmField = "epoch"
)

// Output patterns
const (
	RpmOutputTab     rpmOutputPattern = "tab"
	RpmOutputNewLine rpmOutputPattern = "newline"
)

type (
	// Rpm は、RPMの情報を持つ構造体です.
	Rpm struct {
		name       string
		version    string
		epoch      string
		rel        string
		arch       string
		outPattern rpmOutputPattern
	}
)

func (r Rpm) String() string {
	switch r.outPattern {
	case RpmOutputTab:
		return fmt.Sprintf("[name] %s\t[version] %s\t[rel] %s\t[arch] %s", r.name, r.version, r.rel, r.arch)
	case RpmOutputNewLine:
		return fmt.Sprintf("name   : %s\nversion: %s\nrel    : %s\narch   : %s", r.name, r.version, r.rel, r.arch)
	}

	return ""
}

// Parse は、指定されたファイル名を rpm パッケージ として解析します.
//
// 以下を参考にしました。
//   -https://github.com/rpm-software-management/yum/blob/f8616a2d6e22705371fe6ba47597238d3d1dc2f1/rpmUtils/miscutils.py#L301
func Parse(filename string) *Rpm {
	if filename[len(filename)-4:] == ".rpm" {
		filename = filename[:len(filename)-4]
	}

	var arch string
	archIndex := strings.LastIndex(filename, ".")
	arch = filename[archIndex+1:]

	var rel string
	relIndex := strings.LastIndex(filename[:archIndex], "-")
	rel = filename[relIndex+1 : archIndex]

	var version string
	verIndex := strings.LastIndex(filename[:relIndex], "-")
	version = filename[verIndex+1 : relIndex]

	var epoch string
	epochIndex := strings.Index(filename, ":")
	if epochIndex != -1 {
		epoch = filename[:epochIndex]
	}

	var name string
	name = filename[epochIndex+1 : verIndex]

	return &Rpm{
		name:       name,
		version:    version,
		rel:        rel,
		epoch:      epoch,
		arch:       arch,
		outPattern: RpmOutputTab,
	}
}

// SetOutputPattern は、出力方法を指定します.
func (r *Rpm) SetOutputPattern(p rpmOutputPattern) {
	r.outPattern = p
}

// Get は、指定したフィールドの値を取得します.
//
// 指定できるのは以下のものです.
//   - name
//   - version
//   - rel
//   - arch
//   - epoch
//
// 上記意外を指定した場合、空文字を返します.
func (r *Rpm) Get(field string) string {

	f := rpmField(field)
	switch f {
	case RpmName:
		return r.name
	case RpmVersion:
		return r.version
	case RpmRel:
		return r.rel
	case RpmArch:
		return r.arch
	case RpmEpoch:
		return r.epoch
	}

	return ""
}
