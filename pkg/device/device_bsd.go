//go:build darwin || freebsd

package device

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

var (
	errInvalidCommandOutput = errors.New("invalid command output")
	errFailedToExecute      = errors.New("failed to run")
)

type BSDKernelDevice struct {
	created bool
	index   int
	logger  *zap.Logger
}

func addressFamily(ip net.IPNet) string {
	isV4 := ip.IP.To4() != nil
	if isV4 {
		return "inet"
	}

	return "inet6"
}

func NewKernelDevice(name string) (*BSDKernelDevice, error) {
	if _, err := run("ifconfig", "wg", "create", "name", name); err != nil {
		return nil, err
	}

	return FindKernelDevice(name)
}

func FindKernelDevice(name string) (*BSDKernelDevice, error) {
	i, err := net.InterfaceByName(name)
	if err != nil {
		return nil, err
	}

	return &BSDKernelDevice{
		created: false,
		index:   i.Index,
		logger: zap.L().Named("dev").With(
			zap.String("dev", name),
			zap.String("type", "kernel")),
	}, nil
}

func (d *BSDKernelDevice) Name() string {
	i, err := net.InterfaceByIndex(d.index)
	if err != nil {
		panic(err)
	}

	return i.Name
}

func (d *BSDKernelDevice) Close() error {
	d.logger.Debug("Deleting kernel device")

	_, err := run("ifconfig", d.Name(), "destroy")
	return err
}

func (d *BSDKernelDevice) AddAddress(ip net.IPNet) error {
	d.logger.Debug("Add address", zap.String("addr", ip.String()))

	args := []string{"ifconfig", d.Name(), addressFamily(ip), ip.String()}
	if isV4 := ip.IP.To4() != nil; isV4 {
		// Darwins utun interfaces are L3 point-to-point links
		// which require a destination address for IPv4
		if d.Flags()&net.FlagPointToPoint != 0 {
			args = append(args, ip.IP.String())
		}
	}

	args = append(args, "alias")

	_, err := run(args...)
	return err
}

func (d *BSDKernelDevice) DeleteAddress(ip net.IPNet) error {
	d.logger.Debug("Delete address", zap.String("addr", ip.String()))

	_, err := run("ifconfig", d.Name(), addressFamily(ip), ip.String(), "-alias")
	return err
}

func (d *BSDKernelDevice) Index() int {
	return d.index
}

func (d *BSDKernelDevice) Flags() net.Flags {
	i, err := net.InterfaceByIndex(d.index)
	if err != nil {
		panic(err)
	}

	return i.Flags
}

var mtuRegex = regexp.MustCompile(`(?m)mtu (\d+)`)

func (d *BSDKernelDevice) MTU() int {
	out, err := run("ifconfig", d.Name())
	if err != nil {
		panic(err)
	}

	mtuStrs := mtuRegex.FindStringSubmatch(out)
	if len(mtuStrs) < 2 {
		panic("no MTU found")
	}

	mtu, err := strconv.Atoi(mtuStrs[1])
	if err != nil {
		panic(err)
	}

	return mtu
}

func (d *BSDKernelDevice) SetMTU(mtu int) error {
	d.logger.Debug("Set link MTU", zap.Int("mtu", mtu))

	_, err := run("ifconfig", d.Name(), "mtu", fmt.Sprint(mtu))
	return err
}

func (d *BSDKernelDevice) SetUp() error {
	d.logger.Debug("Set link up")

	i, err := net.InterfaceByIndex(d.index)
	if err != nil {
		return err
	}

	_, err = run("ifconfig", i.Name, "up")
	return err
}

func (d *BSDKernelDevice) SetDown() error {
	d.logger.Debug("Set link down")

	i, err := net.InterfaceByIndex(d.index)
	if err != nil {
		return err
	}

	_, err = run("ifconfig", i.Name, "down")
	return err
}

func Table(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}

	return i, nil
}

func DetectMTU(ip net.IP, fwmark int) (int, error) {
	return getRouteMTU(ip)
}

func DetectDefaultMTU(fwmark int) (int, error) {
	return getRouteMTU(nil)
}

func getRouteMTU(ip net.IP) (int, error) {
	netw := "default"
	if ip != nil {
		netw = ip.String()
	}

	out, err := run("route", "get", netw)
	if err != nil {
		return -1, fmt.Errorf("failed to lookup route: %w", err)
	}

	out = strings.TrimSpace(out)
	lines := strings.Split(out, "\n")
	lastLine := lines[len(lines)-1]
	fields := strings.Fields(lastLine)

	if len(fields) < 7 {
		return -1, fmt.Errorf("%w: %s", errInvalidCommandOutput, lastLine)
	}

	return strconv.Atoi(fields[6])
}

func run(args ...string) (string, error) {
	//#nosec G204 -- Command is always hardcoded
	cmd := exec.Command(args[0], args[1:]...)

	out, err := cmd.CombinedOutput()
	outStr := string(out)
	if err != nil {
		return "", fmt.Errorf("%w: %s\n%s", errFailedToExecute, strings.Join(args, " "), outStr)
	}

	return outStr, nil
}
