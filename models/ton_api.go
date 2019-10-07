package models

import "math/big"

// This class is a base class for all TDLib interface classes.
type TonAPI struct {

}

type AccountAddress struct {
	AccountAddress string
}

type Bip39Hints struct {
	Words []string
}

type Config struct {
	Config string
	blockChainName string
	useCallbacksForNetwork bool
	ignoreCache bool
}

type Error struct {
	Code int
	Message string
}

type ExportedEncryptedKey struct {
	Data []byte
}



type ExportedKey struct {
	wordList []string
}

type ExportedPemKey struct {
	Pem string
}

type InputKey struct {
	// TODO: reverse Java code to understand key type
	Key string

	LocalPassword []byte
}

type Key struct {
	PublicKey string
	Secret []byte
}

type KeyStoreType struct {
	// empty
}

type KeyStoreTypeDirectory struct {
	Directory string
}

type KeyStoreTypeInMemory struct {
	// empty
}

type LogStream struct {
	// empty
}


type LogStreamDefault struct {
	// empty
}

type LogStreamFile struct {
	Path string
	MaxFileSize big.Int
}

type LogStreamEmpty struct {
	// empty
}

type LogTags struct {
	Tags []string
}

type LogVerbosityLevel struct {
	VerbosityLevel int
}

type Ok struct {
	// empty
}

type Options struct {
	Config Config
	KeyStoreType KeyStoreType
}


type SendGramsResult struct {
	SentUntil big.Int
}

type UnpackedAccountAddress struct {
	WorkChainID int
	Bounceble bool
	TestNet bool
	Addr []byte
}

type UpdateSendLiteServerQuery struct {

}

type GenericAccountState struct {

}

type GenericAccountStateRAW struct {

}

type GenericAccountStateTestWallet struct {

}

type GenericAccountStateWallet struct {

}

type GenericAccountStateTestGiver struct {

}

type GenericAccountStateUninited struct {

}

type InternalTransactionID struct {

}

type RAWAccountState struct {

}

type RAWInitialAccountState struct {

}

type RAWMessage struct {

}

type RAWTransaction struct {

}

type TestGiverAccountState struct {

}

type TestWalletAccount struct {

}

type TestWalletInitialAccountState struct {

}

type UninitedAccountState struct {

}

type WalletAccountState struct {

}

type WalletInitialAccountState struct {

}

type AddLogMessage struct {

}

type ChangeLocalPassword struct {

}

type Close struct {

}

type CreateNewKey struct {

}

type DeleteKey struct {

}

type ExportEncryptedKey struct {

}

type ExportKey struct {

}

type ExportPemKey struct {

}

type GenericGetAccountState struct {

}

type GenericSendGrams struct {

}

type GetBip39Hints struct {

}

type GetLogStream struct {

}

type GetLogTagVerbosityLevel struct {

}

type GetLogsTags struct {

}

type GetLogVerbosityLevel struct {

}

type ImportEncryptedKey struct {

}

type ImportKey struct {

}

type Init struct {

}

type OnLiteServerQueryError struct {

}

type OnLiteServerQueryResult struct {

}

type OptionsSetConfig struct {

}

type PackAccountAddress struct {

}

type RAWGetAccountAddress struct {

}

type RAWGetAccountState struct {

}

type RAWGetTransactions struct {

}

type RAWSendMessage struct {

}

type RunTests struct {

}

type SetLogStream struct {

}

type SetLogTagVerbosityLevel struct {

}

type SetLogVerbosityLevel struct {

}

type TestGiverGetAccountAddress struct {

}

type TestGiverGetAccountState struct {

}

type TestGiverSendGrams struct {

}

type TestWalletGetAccountAddress struct {

}

type TestWalletGetAccountState struct {

}

type TestWalletInit struct {

}

type TestWalletSendGrams struct {

}

type UnpackAccountAddress struct {

}

type WalletGetAccountAddress struct {

}

type WalletGetAccountState struct {

}





