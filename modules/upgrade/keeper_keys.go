package upgrade

import (
	"bytes"
	"fmt"
	"strconv"
)

var (
	appVersionKey  = "v/%s/%s"    // v/<protocol.version>/<proposalID>
	proposalIDKey  = "p/%s"    // p/<proposalId>
	successAppVersionKey = "success/%s"    // h/<protocol.version>
	signalKey      = "s/%s/%s" // s/<protocol.version>/<switchVoterAddress>
)

func GetAppVersionKey(versionID uint64, proposalID uint64) []byte {
	return []byte(fmt.Sprintf(appVersionKey, UintToHexString(versionID), UintToHexString(proposalID)))
}

func GetSuccessAppVersionKey(versionID uint64) []byte {
	return []byte(fmt.Sprintf(successAppVersionKey, UintToHexString(versionID)))
}

func GetProposalIDKey(proposalID uint64) []byte {
	return []byte(fmt.Sprintf(proposalIDKey, UintToHexString(proposalID)))
}

func GetSignalKey(versionID uint64, switchVoterAddr string) []byte {
	return []byte(fmt.Sprintf(signalKey, UintToHexString(versionID), switchVoterAddr))
}

func GetPrefixSignalKey(versionID uint64) []byte {
	return []byte(fmt.Sprintf(signalKey, UintToHexString(versionID)))
}

func IntToHexString(i int64) string {
	hex := strconv.FormatInt(i, 16)
	var stringBuild bytes.Buffer
	for i := 0; i < 16-len(hex); i++ {
		stringBuild.Write([]byte("0"))
	}
	stringBuild.Write([]byte(hex))
	return stringBuild.String()
}
func UintToHexString(i uint64) string {
	hex := strconv.FormatUint(i, 16)
	var stringBuild bytes.Buffer
	for i := 0; i < 16-len(hex); i++ {
		stringBuild.Write([]byte("0"))
	}
	stringBuild.Write([]byte(hex))
	return stringBuild.String()
}
