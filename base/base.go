// *********************************************************************************************************************
// ***                                        CONFIDENTIAL --- CUSTOM STUDIOS                                        ***
// *********************************************************************************************************************
// * Auth: ColeCai                                                                                                     *
// * Date: 2023/11/17 21:53:46                                                                                         *
// * Proj: encoding                                                                                                    *
// * Pack: base                                                                                                        *
// * File: base.go                                                                                                     *
// *-------------------------------------------------------------------------------------------------------------------*
// * Overviews:                                                                                                        *
// *-------------------------------------------------------------------------------------------------------------------*
// * Functions:                                                                                                        *
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - *

package base

import "unsafe"

type IEncoding interface {
	Encode(src []byte) ([]byte, error)
	Decode(src []byte) ([]byte, error)
}

func StringToBytes(src string) []byte {
	return unsafe.Slice(unsafe.StringData(src), len(src))
}

func BytesToString(src []byte) string {
	return unsafe.String(&src[0], len(src))
}
