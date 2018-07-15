package account

// AccountClient implements the battlenet.Client interface.
type AccountClient struct {
	userAgent string
	client    *http.Client
	locale    locale.Locale
	token     string
}

// New creates a new AccountClient. Passing different interface types can cause
// different behaviors. See function definiton for more details.
func New(args ...interface{}) (c *AccountClient, err error) {
	c = &AccountClient{
		userAgent: "GoBattleNetAccount/" + battlenet.ClientVersion,
		client:    &http.Client{Timeout: (10 * time.Second)},
		locale:    locale.AmericanEnglish,
		token:     "",
	}

	if nil == args {
		return c, nil
	}

	for _, arg := range args {
		switch t := arg.(type) {
		case string:
			c.token = t
			break
		case *http.Client:
			c.client = t
			break
		case locale.Locale:
			c.locale = t
			break
		case BNetSettings:
			c.client = t.Client
			c.locale = t.Locale
			c.token = t.Token
			break
		default:
			return nil, errors.New(fmt.Sprintf("Type %v is not supported.", t))
		}
	}
	return c, nil
}

// Locale gets the client's locale.
func (c *AccountClient) Locale() locale.Locale {
	return c.locale
}

// SetLocale sets the client's locale.
func (c *AccountClient) SetLocale(locale locale.Locale) {
	c.locale = locale
}

// SetKey sets the client's key.
func (c *AccountClient) SetKey(token string) {
	c.token = token
}

// UserAgent returns the client's user agent.
func (c *AccountClient) UserAgent() string {
	return c.userAgent
}

// Convert an HTTP response from a given URL to the supplied interface.
// This function expects the body to contain the associated JSON response
// from the given URL and will return an error if it fails to properly unmarshal.
func (c *AccountClient) get(url string, v interface{}) error {
	if nil == v {
		return errors.New("No interface supplied.")
	}

	response, err := http.Get(url)
	if nil != err {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err
	}

	err = json.Unmarshal([]byte(body), &v)
	if nil != err {
		return err
	}

	return nil
}
