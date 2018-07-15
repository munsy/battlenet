package battlenet

import (
	"errors"
	"fmt"

	"github.com/munsy/gobattlenet/account"
	"github.com/munsy/gobattlenet/d3"
	"github.com/munsy/gobattlenet/sc2"
	"github.com/munsy/gobattlenet/wow"
)

type ClientType int

const (
	Account ClientType = iota
	Wow
	Sc2
	D3
)

func New(t ClientType, args ...interface{}) (*Client, error) {
	var c *Client
	var err error

	switch t {
	case Account:
		c, err = account.New(args)
		break
	case D3:
		c, err = d3.New(args), nil
		break
	case Sc2:
		c, err = sc2.New(args), nil
		break
	case Wow:
		c, err = wow.New(args), nil
		break
	default:
		c = nil
		err = errors.New(fmt.Sprintf("%v is not supported.", t))
		break
	}

	return c, err
}
