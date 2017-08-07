// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: account.go

package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Account) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *Account) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	/* Struct fall back. type=objects.GrapheneID kind=struct */
	buf.WriteString(`{"id":`)
	err = buf.Encode(&j.ID)
	if err != nil {
		return err
	}
	buf.WriteString(`,"name":`)
	fflib.WriteJsonString(buf, string(j.Name))
	/* Struct fall back. type=objects.GrapheneID kind=struct */
	buf.WriteString(`,"statistics":`)
	err = buf.Encode(&j.Statistics)
	if err != nil {
		return err
	}
	buf.WriteString(`,"membership_expiration_date":`)

	{

		obj, err = j.MembershipExpirationDate.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"network_fee_percentage":`)
	fflib.FormatBits2(buf, uint64(j.NetworkFeePercentage), 10, j.NetworkFeePercentage < 0)
	buf.WriteString(`,"lifetime_referrer_fee_percentage":`)
	fflib.FormatBits2(buf, uint64(j.LifetimeReferrerFeePercentage), 10, j.LifetimeReferrerFeePercentage < 0)
	buf.WriteString(`,"referrer_rewards_percentage":`)
	fflib.FormatBits2(buf, uint64(j.ReferrerRewardsPercentage), 10, j.ReferrerRewardsPercentage < 0)
	buf.WriteString(`,"top_n_control_flags":`)
	fflib.FormatBits2(buf, uint64(j.TopNControlFlags), 10, j.TopNControlFlags < 0)
	buf.WriteString(`,"whitelisting_accounts":`)
	if j.WhitelistingAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.WhitelistingAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}
			fflib.WriteJsonString(buf, string(v))
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"blacklisting_accounts":`)
	if j.BlacklistingAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.BlacklistingAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}
			fflib.WriteJsonString(buf, string(v))
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"whitelisted_accounts":`)
	if j.WhitelistedAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.WhitelistedAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}
			fflib.WriteJsonString(buf, string(v))
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"blacklisted_accounts":`)
	if j.BlacklistedAccounts != nil {
		buf.WriteString(`[`)
		for i, v := range j.BlacklistedAccounts {
			if i != 0 {
				buf.WriteString(`,`)
			}
			fflib.WriteJsonString(buf, string(v))
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	/* Struct fall back. type=objects.AccountOptions kind=struct */
	buf.WriteString(`,"options":`)
	err = buf.Encode(&j.Options)
	if err != nil {
		return err
	}
	/* Struct fall back. type=objects.GrapheneID kind=struct */
	buf.WriteString(`,"registrar":`)
	err = buf.Encode(&j.Registrar)
	if err != nil {
		return err
	}
	/* Struct fall back. type=objects.GrapheneID kind=struct */
	buf.WriteString(`,"referrer":`)
	err = buf.Encode(&j.Referrer)
	if err != nil {
		return err
	}
	/* Struct fall back. type=objects.GrapheneID kind=struct */
	buf.WriteString(`,"lifetime_referrer":`)
	err = buf.Encode(&j.LifetimeReferrer)
	if err != nil {
		return err
	}
	/* Struct fall back. type=objects.GrapheneID kind=struct */
	buf.WriteString(`,"cashback_vb":`)
	err = buf.Encode(&j.CashbackVB)
	if err != nil {
		return err
	}
	/* Struct fall back. type=objects.Authority kind=struct */
	buf.WriteString(`,"owner":`)
	err = buf.Encode(&j.Owner)
	if err != nil {
		return err
	}
	/* Struct fall back. type=objects.Authority kind=struct */
	buf.WriteString(`,"active":`)
	err = buf.Encode(&j.Active)
	if err != nil {
		return err
	}
	buf.WriteString(`,"owner_special_authority":`)
	if j.OwnerSpecialAuthority != nil {
		buf.WriteString(`[`)
		for i, v := range j.OwnerSpecialAuthority {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Interface types must use runtime reflection. type=interface {} kind=interface */
			err = buf.Encode(v)
			if err != nil {
				return err
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"active_special_authority":`)
	if j.ActiveSpecialAuthority != nil {
		buf.WriteString(`[`)
		for i, v := range j.ActiveSpecialAuthority {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Interface types must use runtime reflection. type=interface {} kind=interface */
			err = buf.Encode(v)
			if err != nil {
				return err
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAccountbase = iota
	ffjtAccountnosuchkey

	ffjtAccountID

	ffjtAccountName

	ffjtAccountStatistics

	ffjtAccountMembershipExpirationDate

	ffjtAccountNetworkFeePercentage

	ffjtAccountLifetimeReferrerFeePercentage

	ffjtAccountReferrerRewardsPercentage

	ffjtAccountTopNControlFlags

	ffjtAccountWhitelistingAccounts

	ffjtAccountBlacklistingAccounts

	ffjtAccountWhitelistedAccounts

	ffjtAccountBlacklistedAccounts

	ffjtAccountOptions

	ffjtAccountRegistrar

	ffjtAccountReferrer

	ffjtAccountLifetimeReferrer

	ffjtAccountCashbackVB

	ffjtAccountOwner

	ffjtAccountActive

	ffjtAccountOwnerSpecialAuthority

	ffjtAccountActiveSpecialAuthority
)

var ffjKeyAccountID = []byte("id")

var ffjKeyAccountName = []byte("name")

var ffjKeyAccountStatistics = []byte("statistics")

var ffjKeyAccountMembershipExpirationDate = []byte("membership_expiration_date")

var ffjKeyAccountNetworkFeePercentage = []byte("network_fee_percentage")

var ffjKeyAccountLifetimeReferrerFeePercentage = []byte("lifetime_referrer_fee_percentage")

var ffjKeyAccountReferrerRewardsPercentage = []byte("referrer_rewards_percentage")

var ffjKeyAccountTopNControlFlags = []byte("top_n_control_flags")

var ffjKeyAccountWhitelistingAccounts = []byte("whitelisting_accounts")

var ffjKeyAccountBlacklistingAccounts = []byte("blacklisting_accounts")

var ffjKeyAccountWhitelistedAccounts = []byte("whitelisted_accounts")

var ffjKeyAccountBlacklistedAccounts = []byte("blacklisted_accounts")

var ffjKeyAccountOptions = []byte("options")

var ffjKeyAccountRegistrar = []byte("registrar")

var ffjKeyAccountReferrer = []byte("referrer")

var ffjKeyAccountLifetimeReferrer = []byte("lifetime_referrer")

var ffjKeyAccountCashbackVB = []byte("cashback_vb")

var ffjKeyAccountOwner = []byte("owner")

var ffjKeyAccountActive = []byte("active")

var ffjKeyAccountOwnerSpecialAuthority = []byte("owner_special_authority")

var ffjKeyAccountActiveSpecialAuthority = []byte("active_special_authority")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Account) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Account) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAccountbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAccountnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAccountActive, kn) {
						currentKey = ffjtAccountActive
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountActiveSpecialAuthority, kn) {
						currentKey = ffjtAccountActiveSpecialAuthority
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyAccountBlacklistingAccounts, kn) {
						currentKey = ffjtAccountBlacklistingAccounts
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountBlacklistedAccounts, kn) {
						currentKey = ffjtAccountBlacklistedAccounts
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyAccountCashbackVB, kn) {
						currentKey = ffjtAccountCashbackVB
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyAccountID, kn) {
						currentKey = ffjtAccountID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyAccountLifetimeReferrerFeePercentage, kn) {
						currentKey = ffjtAccountLifetimeReferrerFeePercentage
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountLifetimeReferrer, kn) {
						currentKey = ffjtAccountLifetimeReferrer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyAccountMembershipExpirationDate, kn) {
						currentKey = ffjtAccountMembershipExpirationDate
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyAccountName, kn) {
						currentKey = ffjtAccountName
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountNetworkFeePercentage, kn) {
						currentKey = ffjtAccountNetworkFeePercentage
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyAccountOptions, kn) {
						currentKey = ffjtAccountOptions
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountOwner, kn) {
						currentKey = ffjtAccountOwner
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountOwnerSpecialAuthority, kn) {
						currentKey = ffjtAccountOwnerSpecialAuthority
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyAccountReferrerRewardsPercentage, kn) {
						currentKey = ffjtAccountReferrerRewardsPercentage
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountRegistrar, kn) {
						currentKey = ffjtAccountRegistrar
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountReferrer, kn) {
						currentKey = ffjtAccountReferrer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyAccountStatistics, kn) {
						currentKey = ffjtAccountStatistics
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyAccountTopNControlFlags, kn) {
						currentKey = ffjtAccountTopNControlFlags
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyAccountWhitelistingAccounts, kn) {
						currentKey = ffjtAccountWhitelistingAccounts
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAccountWhitelistedAccounts, kn) {
						currentKey = ffjtAccountWhitelistedAccounts
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyAccountActiveSpecialAuthority, kn) {
					currentKey = ffjtAccountActiveSpecialAuthority
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountOwnerSpecialAuthority, kn) {
					currentKey = ffjtAccountOwnerSpecialAuthority
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountActive, kn) {
					currentKey = ffjtAccountActive
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountOwner, kn) {
					currentKey = ffjtAccountOwner
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountCashbackVB, kn) {
					currentKey = ffjtAccountCashbackVB
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAccountLifetimeReferrer, kn) {
					currentKey = ffjtAccountLifetimeReferrer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountReferrer, kn) {
					currentKey = ffjtAccountReferrer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountRegistrar, kn) {
					currentKey = ffjtAccountRegistrar
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountOptions, kn) {
					currentKey = ffjtAccountOptions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountBlacklistedAccounts, kn) {
					currentKey = ffjtAccountBlacklistedAccounts
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountWhitelistedAccounts, kn) {
					currentKey = ffjtAccountWhitelistedAccounts
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountBlacklistingAccounts, kn) {
					currentKey = ffjtAccountBlacklistingAccounts
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountWhitelistingAccounts, kn) {
					currentKey = ffjtAccountWhitelistingAccounts
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountTopNControlFlags, kn) {
					currentKey = ffjtAccountTopNControlFlags
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountReferrerRewardsPercentage, kn) {
					currentKey = ffjtAccountReferrerRewardsPercentage
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAccountLifetimeReferrerFeePercentage, kn) {
					currentKey = ffjtAccountLifetimeReferrerFeePercentage
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountNetworkFeePercentage, kn) {
					currentKey = ffjtAccountNetworkFeePercentage
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountMembershipExpirationDate, kn) {
					currentKey = ffjtAccountMembershipExpirationDate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAccountStatistics, kn) {
					currentKey = ffjtAccountStatistics
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountName, kn) {
					currentKey = ffjtAccountName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAccountID, kn) {
					currentKey = ffjtAccountID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAccountnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAccountID:
					goto handle_ID

				case ffjtAccountName:
					goto handle_Name

				case ffjtAccountStatistics:
					goto handle_Statistics

				case ffjtAccountMembershipExpirationDate:
					goto handle_MembershipExpirationDate

				case ffjtAccountNetworkFeePercentage:
					goto handle_NetworkFeePercentage

				case ffjtAccountLifetimeReferrerFeePercentage:
					goto handle_LifetimeReferrerFeePercentage

				case ffjtAccountReferrerRewardsPercentage:
					goto handle_ReferrerRewardsPercentage

				case ffjtAccountTopNControlFlags:
					goto handle_TopNControlFlags

				case ffjtAccountWhitelistingAccounts:
					goto handle_WhitelistingAccounts

				case ffjtAccountBlacklistingAccounts:
					goto handle_BlacklistingAccounts

				case ffjtAccountWhitelistedAccounts:
					goto handle_WhitelistedAccounts

				case ffjtAccountBlacklistedAccounts:
					goto handle_BlacklistedAccounts

				case ffjtAccountOptions:
					goto handle_Options

				case ffjtAccountRegistrar:
					goto handle_Registrar

				case ffjtAccountReferrer:
					goto handle_Referrer

				case ffjtAccountLifetimeReferrer:
					goto handle_LifetimeReferrer

				case ffjtAccountCashbackVB:
					goto handle_CashbackVB

				case ffjtAccountOwner:
					goto handle_Owner

				case ffjtAccountActive:
					goto handle_Active

				case ffjtAccountOwnerSpecialAuthority:
					goto handle_OwnerSpecialAuthority

				case ffjtAccountActiveSpecialAuthority:
					goto handle_ActiveSpecialAuthority

				case ffjtAccountnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_ID:

	/* handler: j.ID type=objects.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = j.ID.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Name:

	/* handler: j.Name type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Name = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Statistics:

	/* handler: j.Statistics type=objects.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = j.Statistics.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MembershipExpirationDate:

	/* handler: j.MembershipExpirationDate type=objects.RFC3339Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = j.MembershipExpirationDate.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NetworkFeePercentage:

	/* handler: j.NetworkFeePercentage type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.NetworkFeePercentage = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LifetimeReferrerFeePercentage:

	/* handler: j.LifetimeReferrerFeePercentage type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.LifetimeReferrerFeePercentage = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ReferrerRewardsPercentage:

	/* handler: j.ReferrerRewardsPercentage type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.ReferrerRewardsPercentage = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TopNControlFlags:

	/* handler: j.TopNControlFlags type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.TopNControlFlags = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WhitelistingAccounts:

	/* handler: j.WhitelistingAccounts type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.WhitelistingAccounts = nil
		} else {

			j.WhitelistingAccounts = []string{}

			wantVal := true

			for {

				var tmpJWhitelistingAccounts string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJWhitelistingAccounts type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJWhitelistingAccounts = string(string(outBuf))

					}
				}

				j.WhitelistingAccounts = append(j.WhitelistingAccounts, tmpJWhitelistingAccounts)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlacklistingAccounts:

	/* handler: j.BlacklistingAccounts type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BlacklistingAccounts = nil
		} else {

			j.BlacklistingAccounts = []string{}

			wantVal := true

			for {

				var tmpJBlacklistingAccounts string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBlacklistingAccounts type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBlacklistingAccounts = string(string(outBuf))

					}
				}

				j.BlacklistingAccounts = append(j.BlacklistingAccounts, tmpJBlacklistingAccounts)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WhitelistedAccounts:

	/* handler: j.WhitelistedAccounts type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.WhitelistedAccounts = nil
		} else {

			j.WhitelistedAccounts = []string{}

			wantVal := true

			for {

				var tmpJWhitelistedAccounts string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJWhitelistedAccounts type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJWhitelistedAccounts = string(string(outBuf))

					}
				}

				j.WhitelistedAccounts = append(j.WhitelistedAccounts, tmpJWhitelistedAccounts)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlacklistedAccounts:

	/* handler: j.BlacklistedAccounts type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BlacklistedAccounts = nil
		} else {

			j.BlacklistedAccounts = []string{}

			wantVal := true

			for {

				var tmpJBlacklistedAccounts string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBlacklistedAccounts type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBlacklistedAccounts = string(string(outBuf))

					}
				}

				j.BlacklistedAccounts = append(j.BlacklistedAccounts, tmpJBlacklistedAccounts)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Options:

	/* handler: j.Options type=objects.AccountOptions kind=struct quoted=false*/

	{
		/* Falling back. type=objects.AccountOptions kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Options)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Registrar:

	/* handler: j.Registrar type=objects.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = j.Registrar.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Referrer:

	/* handler: j.Referrer type=objects.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = j.Referrer.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LifetimeReferrer:

	/* handler: j.LifetimeReferrer type=objects.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = j.LifetimeReferrer.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CashbackVB:

	/* handler: j.CashbackVB type=objects.GrapheneID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = j.CashbackVB.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Owner:

	/* handler: j.Owner type=objects.Authority kind=struct quoted=false*/

	{
		/* Falling back. type=objects.Authority kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Owner)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Active:

	/* handler: j.Active type=objects.Authority kind=struct quoted=false*/

	{
		/* Falling back. type=objects.Authority kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Active)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_OwnerSpecialAuthority:

	/* handler: j.OwnerSpecialAuthority type=[]interface {} kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.OwnerSpecialAuthority = nil
		} else {

			j.OwnerSpecialAuthority = []interface{}{}

			wantVal := true

			for {

				var tmpJOwnerSpecialAuthority interface{}

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJOwnerSpecialAuthority type=interface {} kind=interface quoted=false*/

				{
					/* Falling back. type=interface {} kind=interface */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJOwnerSpecialAuthority)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.OwnerSpecialAuthority = append(j.OwnerSpecialAuthority, tmpJOwnerSpecialAuthority)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ActiveSpecialAuthority:

	/* handler: j.ActiveSpecialAuthority type=[]interface {} kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.ActiveSpecialAuthority = nil
		} else {

			j.ActiveSpecialAuthority = []interface{}{}

			wantVal := true

			for {

				var tmpJActiveSpecialAuthority interface{}

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJActiveSpecialAuthority type=interface {} kind=interface quoted=false*/

				{
					/* Falling back. type=interface {} kind=interface */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJActiveSpecialAuthority)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.ActiveSpecialAuthority = append(j.ActiveSpecialAuthority, tmpJActiveSpecialAuthority)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
