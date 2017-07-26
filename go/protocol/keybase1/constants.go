// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/constants.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type StatusCode int

const (
	StatusCode_SCOk                       StatusCode = 0
	StatusCode_SCInputError               StatusCode = 100
	StatusCode_SCLoginRequired            StatusCode = 201
	StatusCode_SCBadSession               StatusCode = 202
	StatusCode_SCBadLoginUserNotFound     StatusCode = 203
	StatusCode_SCBadLoginPassword         StatusCode = 204
	StatusCode_SCNotFound                 StatusCode = 205
	StatusCode_SCThrottleControl          StatusCode = 210
	StatusCode_SCDeleted                  StatusCode = 216
	StatusCode_SCGeneric                  StatusCode = 218
	StatusCode_SCAlreadyLoggedIn          StatusCode = 235
	StatusCode_SCExists                   StatusCode = 230
	StatusCode_SCCanceled                 StatusCode = 237
	StatusCode_SCInputCanceled            StatusCode = 239
	StatusCode_SCReloginRequired          StatusCode = 274
	StatusCode_SCResolutionFailed         StatusCode = 275
	StatusCode_SCProfileNotPublic         StatusCode = 276
	StatusCode_SCIdentifyFailed           StatusCode = 277
	StatusCode_SCTrackingBroke            StatusCode = 278
	StatusCode_SCWrongCryptoFormat        StatusCode = 279
	StatusCode_SCDecryptionError          StatusCode = 280
	StatusCode_SCInvalidAddress           StatusCode = 281
	StatusCode_SCNoSession                StatusCode = 283
	StatusCode_SCAccountReset             StatusCode = 290
	StatusCode_SCBadEmail                 StatusCode = 472
	StatusCode_SCBadSignupUsernameTaken   StatusCode = 701
	StatusCode_SCBadInvitationCode        StatusCode = 707
	StatusCode_SCMissingResult            StatusCode = 801
	StatusCode_SCKeyNotFound              StatusCode = 901
	StatusCode_SCKeyCorrupted             StatusCode = 905
	StatusCode_SCKeyInUse                 StatusCode = 907
	StatusCode_SCKeyBadGen                StatusCode = 913
	StatusCode_SCKeyNoSecret              StatusCode = 914
	StatusCode_SCKeyBadUIDs               StatusCode = 915
	StatusCode_SCKeyNoActive              StatusCode = 916
	StatusCode_SCKeyNoSig                 StatusCode = 917
	StatusCode_SCKeyBadSig                StatusCode = 918
	StatusCode_SCKeyBadEldest             StatusCode = 919
	StatusCode_SCKeyNoEldest              StatusCode = 920
	StatusCode_SCKeyDuplicateUpdate       StatusCode = 921
	StatusCode_SCSibkeyAlreadyExists      StatusCode = 922
	StatusCode_SCDecryptionKeyNotFound    StatusCode = 924
	StatusCode_SCKeyNoPGPEncryption       StatusCode = 927
	StatusCode_SCKeyNoNaClEncryption      StatusCode = 928
	StatusCode_SCKeySyncedPGPNotFound     StatusCode = 929
	StatusCode_SCKeyNoMatchingGPG         StatusCode = 930
	StatusCode_SCKeyRevoked               StatusCode = 931
	StatusCode_SCBadTrackSession          StatusCode = 1301
	StatusCode_SCDeviceBadName            StatusCode = 1404
	StatusCode_SCDeviceNameInUse          StatusCode = 1408
	StatusCode_SCDeviceNotFound           StatusCode = 1409
	StatusCode_SCDeviceMismatch           StatusCode = 1410
	StatusCode_SCDeviceRequired           StatusCode = 1411
	StatusCode_SCDevicePrevProvisioned    StatusCode = 1413
	StatusCode_SCDeviceNoProvision        StatusCode = 1414
	StatusCode_SCDeviceProvisionViaDevice StatusCode = 1415
	StatusCode_SCStreamExists             StatusCode = 1501
	StatusCode_SCStreamNotFound           StatusCode = 1502
	StatusCode_SCStreamWrongKind          StatusCode = 1503
	StatusCode_SCStreamEOF                StatusCode = 1504
	StatusCode_SCGenericAPIError          StatusCode = 1600
	StatusCode_SCAPINetworkError          StatusCode = 1601
	StatusCode_SCTimeout                  StatusCode = 1602
	StatusCode_SCProofError               StatusCode = 1701
	StatusCode_SCIdentificationExpired    StatusCode = 1702
	StatusCode_SCSelfNotFound             StatusCode = 1703
	StatusCode_SCBadKexPhrase             StatusCode = 1704
	StatusCode_SCNoUIDelegation           StatusCode = 1705
	StatusCode_SCNoUI                     StatusCode = 1706
	StatusCode_SCGPGUnavailable           StatusCode = 1707
	StatusCode_SCInvalidVersionError      StatusCode = 1800
	StatusCode_SCOldVersionError          StatusCode = 1801
	StatusCode_SCInvalidLocationError     StatusCode = 1802
	StatusCode_SCServiceStatusError       StatusCode = 1803
	StatusCode_SCInstallError             StatusCode = 1804
	StatusCode_SCLoginStateTimeout        StatusCode = 2400
	StatusCode_SCChatInternal             StatusCode = 2500
	StatusCode_SCChatRateLimit            StatusCode = 2501
	StatusCode_SCChatConvExists           StatusCode = 2502
	StatusCode_SCChatUnknownTLFID         StatusCode = 2503
	StatusCode_SCChatNotInConv            StatusCode = 2504
	StatusCode_SCChatBadMsg               StatusCode = 2505
	StatusCode_SCChatBroadcast            StatusCode = 2506
	StatusCode_SCChatAlreadySuperseded    StatusCode = 2507
	StatusCode_SCChatAlreadyDeleted       StatusCode = 2508
	StatusCode_SCChatTLFFinalized         StatusCode = 2509
	StatusCode_SCChatCollision            StatusCode = 2510
	StatusCode_SCIdentifySummaryError     StatusCode = 2511
	StatusCode_SCNeedSelfRekey            StatusCode = 2512
	StatusCode_SCNeedOtherRekey           StatusCode = 2513
	StatusCode_SCChatMessageCollision     StatusCode = 2514
	StatusCode_SCChatDuplicateMessage     StatusCode = 2515
	StatusCode_SCChatClientError          StatusCode = 2516
	StatusCode_SCChatNotInTeam            StatusCode = 2517
	StatusCode_SCTeamExists               StatusCode = 2619
	StatusCode_SCTeamReadError            StatusCode = 2623
	StatusCode_SCTeamTarDuplicate         StatusCode = 2663
	StatusCode_SCTeamTarNotFound          StatusCode = 2664
	StatusCode_SCTeamMemberExists         StatusCode = 2665
)

func (o StatusCode) DeepCopy() StatusCode { return o }

var StatusCodeMap = map[string]StatusCode{
	"SCOk":                       0,
	"SCInputError":               100,
	"SCLoginRequired":            201,
	"SCBadSession":               202,
	"SCBadLoginUserNotFound":     203,
	"SCBadLoginPassword":         204,
	"SCNotFound":                 205,
	"SCThrottleControl":          210,
	"SCDeleted":                  216,
	"SCGeneric":                  218,
	"SCAlreadyLoggedIn":          235,
	"SCExists":                   230,
	"SCCanceled":                 237,
	"SCInputCanceled":            239,
	"SCReloginRequired":          274,
	"SCResolutionFailed":         275,
	"SCProfileNotPublic":         276,
	"SCIdentifyFailed":           277,
	"SCTrackingBroke":            278,
	"SCWrongCryptoFormat":        279,
	"SCDecryptionError":          280,
	"SCInvalidAddress":           281,
	"SCNoSession":                283,
	"SCAccountReset":             290,
	"SCBadEmail":                 472,
	"SCBadSignupUsernameTaken":   701,
	"SCBadInvitationCode":        707,
	"SCMissingResult":            801,
	"SCKeyNotFound":              901,
	"SCKeyCorrupted":             905,
	"SCKeyInUse":                 907,
	"SCKeyBadGen":                913,
	"SCKeyNoSecret":              914,
	"SCKeyBadUIDs":               915,
	"SCKeyNoActive":              916,
	"SCKeyNoSig":                 917,
	"SCKeyBadSig":                918,
	"SCKeyBadEldest":             919,
	"SCKeyNoEldest":              920,
	"SCKeyDuplicateUpdate":       921,
	"SCSibkeyAlreadyExists":      922,
	"SCDecryptionKeyNotFound":    924,
	"SCKeyNoPGPEncryption":       927,
	"SCKeyNoNaClEncryption":      928,
	"SCKeySyncedPGPNotFound":     929,
	"SCKeyNoMatchingGPG":         930,
	"SCKeyRevoked":               931,
	"SCBadTrackSession":          1301,
	"SCDeviceBadName":            1404,
	"SCDeviceNameInUse":          1408,
	"SCDeviceNotFound":           1409,
	"SCDeviceMismatch":           1410,
	"SCDeviceRequired":           1411,
	"SCDevicePrevProvisioned":    1413,
	"SCDeviceNoProvision":        1414,
	"SCDeviceProvisionViaDevice": 1415,
	"SCStreamExists":             1501,
	"SCStreamNotFound":           1502,
	"SCStreamWrongKind":          1503,
	"SCStreamEOF":                1504,
	"SCGenericAPIError":          1600,
	"SCAPINetworkError":          1601,
	"SCTimeout":                  1602,
	"SCProofError":               1701,
	"SCIdentificationExpired":    1702,
	"SCSelfNotFound":             1703,
	"SCBadKexPhrase":             1704,
	"SCNoUIDelegation":           1705,
	"SCNoUI":                     1706,
	"SCGPGUnavailable":           1707,
	"SCInvalidVersionError":      1800,
	"SCOldVersionError":          1801,
	"SCInvalidLocationError":     1802,
	"SCServiceStatusError":       1803,
	"SCInstallError":             1804,
	"SCLoginStateTimeout":        2400,
	"SCChatInternal":             2500,
	"SCChatRateLimit":            2501,
	"SCChatConvExists":           2502,
	"SCChatUnknownTLFID":         2503,
	"SCChatNotInConv":            2504,
	"SCChatBadMsg":               2505,
	"SCChatBroadcast":            2506,
	"SCChatAlreadySuperseded":    2507,
	"SCChatAlreadyDeleted":       2508,
	"SCChatTLFFinalized":         2509,
	"SCChatCollision":            2510,
	"SCIdentifySummaryError":     2511,
	"SCNeedSelfRekey":            2512,
	"SCNeedOtherRekey":           2513,
	"SCChatMessageCollision":     2514,
	"SCChatDuplicateMessage":     2515,
	"SCChatClientError":          2516,
	"SCChatNotInTeam":            2517,
	"SCTeamExists":               2619,
	"SCTeamReadError":            2623,
	"SCTeamTarDuplicate":         2663,
	"SCTeamTarNotFound":          2664,
	"SCTeamMemberExists":         2665,
}

var StatusCodeRevMap = map[StatusCode]string{
	0:    "SCOk",
	100:  "SCInputError",
	201:  "SCLoginRequired",
	202:  "SCBadSession",
	203:  "SCBadLoginUserNotFound",
	204:  "SCBadLoginPassword",
	205:  "SCNotFound",
	210:  "SCThrottleControl",
	216:  "SCDeleted",
	218:  "SCGeneric",
	235:  "SCAlreadyLoggedIn",
	230:  "SCExists",
	237:  "SCCanceled",
	239:  "SCInputCanceled",
	274:  "SCReloginRequired",
	275:  "SCResolutionFailed",
	276:  "SCProfileNotPublic",
	277:  "SCIdentifyFailed",
	278:  "SCTrackingBroke",
	279:  "SCWrongCryptoFormat",
	280:  "SCDecryptionError",
	281:  "SCInvalidAddress",
	283:  "SCNoSession",
	290:  "SCAccountReset",
	472:  "SCBadEmail",
	701:  "SCBadSignupUsernameTaken",
	707:  "SCBadInvitationCode",
	801:  "SCMissingResult",
	901:  "SCKeyNotFound",
	905:  "SCKeyCorrupted",
	907:  "SCKeyInUse",
	913:  "SCKeyBadGen",
	914:  "SCKeyNoSecret",
	915:  "SCKeyBadUIDs",
	916:  "SCKeyNoActive",
	917:  "SCKeyNoSig",
	918:  "SCKeyBadSig",
	919:  "SCKeyBadEldest",
	920:  "SCKeyNoEldest",
	921:  "SCKeyDuplicateUpdate",
	922:  "SCSibkeyAlreadyExists",
	924:  "SCDecryptionKeyNotFound",
	927:  "SCKeyNoPGPEncryption",
	928:  "SCKeyNoNaClEncryption",
	929:  "SCKeySyncedPGPNotFound",
	930:  "SCKeyNoMatchingGPG",
	931:  "SCKeyRevoked",
	1301: "SCBadTrackSession",
	1404: "SCDeviceBadName",
	1408: "SCDeviceNameInUse",
	1409: "SCDeviceNotFound",
	1410: "SCDeviceMismatch",
	1411: "SCDeviceRequired",
	1413: "SCDevicePrevProvisioned",
	1414: "SCDeviceNoProvision",
	1415: "SCDeviceProvisionViaDevice",
	1501: "SCStreamExists",
	1502: "SCStreamNotFound",
	1503: "SCStreamWrongKind",
	1504: "SCStreamEOF",
	1600: "SCGenericAPIError",
	1601: "SCAPINetworkError",
	1602: "SCTimeout",
	1701: "SCProofError",
	1702: "SCIdentificationExpired",
	1703: "SCSelfNotFound",
	1704: "SCBadKexPhrase",
	1705: "SCNoUIDelegation",
	1706: "SCNoUI",
	1707: "SCGPGUnavailable",
	1800: "SCInvalidVersionError",
	1801: "SCOldVersionError",
	1802: "SCInvalidLocationError",
	1803: "SCServiceStatusError",
	1804: "SCInstallError",
	2400: "SCLoginStateTimeout",
	2500: "SCChatInternal",
	2501: "SCChatRateLimit",
	2502: "SCChatConvExists",
	2503: "SCChatUnknownTLFID",
	2504: "SCChatNotInConv",
	2505: "SCChatBadMsg",
	2506: "SCChatBroadcast",
	2507: "SCChatAlreadySuperseded",
	2508: "SCChatAlreadyDeleted",
	2509: "SCChatTLFFinalized",
	2510: "SCChatCollision",
	2511: "SCIdentifySummaryError",
	2512: "SCNeedSelfRekey",
	2513: "SCNeedOtherRekey",
	2514: "SCChatMessageCollision",
	2515: "SCChatDuplicateMessage",
	2516: "SCChatClientError",
	2517: "SCChatNotInTeam",
	2619: "SCTeamExists",
	2623: "SCTeamReadError",
	2663: "SCTeamTarDuplicate",
	2664: "SCTeamTarNotFound",
	2665: "SCTeamMemberExists",
}

func (e StatusCode) String() string {
	if v, ok := StatusCodeRevMap[e]; ok {
		return v
	}
	return ""
}

type ConstantsInterface interface {
}

func ConstantsProtocol(i ConstantsInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "keybase.1.constants",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type ConstantsClient struct {
	Cli rpc.GenericClient
}
