
// GENERATED FILE, DO NOT EDIT!

package cmac

type wycheproofVector struct {
	TcId   int
	Flag  string
	Key    string
	Msg    string
	Tag    string
	Result string
}

var wycheproofVectors = []wycheproofVector{

{TcId: 1, Flag: "Pseudorandom", Key: "e34f15c7bd819930fe9d66e0c166e61c", Msg: "", Tag: "d47afca1d857a5933405b1eb7a5cb7af", Result: "valid"},

{TcId: 2, Flag: "Pseudorandom", Key: "e1e726677f4893890f8c027f9d8ef80d", Msg: "3f", Tag: "15f856bbed3b321952a584b3c4437a63", Result: "valid"},

{TcId: 3, Flag: "Pseudorandom", Key: "b151f491c4c006d1f28214aa3da9a985", Msg: "27d9", Tag: "bdbbebac982dd62b9f682618a6a604e9", Result: "valid"},

{TcId: 4, Flag: "Pseudorandom", Key: "c36ff15f72777ee21deec07b63c1a0cd", Msg: "50b428", Tag: "be0c3ede157568af394023eb9a7cc983", Result: "valid"},

{TcId: 5, Flag: "Pseudorandom", Key: "32b9c5c78c3a0689a86052420fa1e8fc", Msg: "0b9262ec", Tag: "57e1506856c55dd32cd9ca821adb6c81", Result: "valid"},

{TcId: 6, Flag: "Pseudorandom", Key: "43151bbaef367277ebfc97509d0aa49c", Msg: "eaa91273e7", Tag: "e01adc3be6a7621824232c4285dd35b9", Result: "valid"},

{TcId: 7, Flag: "Pseudorandom", Key: "481440298525cc261f8159159aedf62d", Msg: "6123c556c5cc", Tag: "a281e0d2d5378dfdcc1310fd9782ca56", Result: "valid"},

{TcId: 8, Flag: "Pseudorandom", Key: "9ca26eb88731efbf7f810d5d95e196ac", Msg: "7e48f06183aa40", Tag: "fc81761f2f7b4ce13b53d36e32677332", Result: "valid"},

{TcId: 9, Flag: "Pseudorandom", Key: "48f0d03e41cc55c4b58f737b5acdea32", Msg: "f4a133aa6d5985a0", Tag: "1f1cd0327c02e6d00086915937dd61d9", Result: "valid"},

{TcId: 10, Flag: "Pseudorandom", Key: "1c958849f31996b28939ce513087d1be", Msg: "b0d2fee11b8e2f86b7", Tag: "555f462151f7dd16de698d639fb26760", Result: "valid"},

{TcId: 11, Flag: "Pseudorandom", Key: "39de0ebea97c09b2301a90009a423253", Msg: "81e5c33b4c620852f044", Tag: "9b004f15b7f6f366374954e64bc58f5f", Result: "valid"},

{TcId: 12, Flag: "Pseudorandom", Key: "91656d8fc0aced60ddb1c4006d0dde53", Msg: "7b3e440fe566790064b2ec", Tag: "76672ed16c29be449e0c80785cc38e89", Result: "valid"},

{TcId: 13, Flag: "Pseudorandom", Key: "af7d5134720b5386158d51ea126e7cf9", Msg: "7cc6fcc925c20f3c83b5567c", Tag: "2dc5c88cf3b80ab6c0199f40be904abc", Result: "valid"},

{TcId: 14, Flag: "Pseudorandom", Key: "4ed56753de6f75a032ebabca3ce27971", Msg: "0c8c0f5619d9f8da5339281285", Tag: "eab4366d97e99a0850f077329ad058c0", Result: "valid"},

{TcId: 15, Flag: "Pseudorandom", Key: "beba50c936b696c15e25046dffb23a64", Msg: "821ea8532fbabffb6e3d212e9b46", Tag: "22f33cab09c173f75d3401fe44efeead", Result: "valid"},

{TcId: 16, Flag: "Pseudorandom", Key: "501d81ebf912ddb87fbe3b7aac1437bc", Msg: "2368e3c3636b5e8e94d2081adbf798", Tag: "aeb784a3825168ddd61f72d0202125e6", Result: "valid"},

{TcId: 17, Flag: "Pseudorandom", Key: "e09eaa5a3f5e56d279d5e7a03373f6ea", Msg: "ef4eab37181f98423e53e947e7050fd0", Tag: "40facf0e2fb51b73a7472681b033d6dc", Result: "valid"},

{TcId: 18, Flag: "Pseudorandom", Key: "831e664c9e3f0c3094c0b27b9d908eb2", Msg: "26603bb76dd0a0180791c4ed4d3b058807", Tag: "a8144c8b24f2aa47d9c160cff4ab1716", Result: "valid"},

{TcId: 19, Flag: "Pseudorandom", Key: "cbffc6c8c7f76f46349c32d666f4efb0", Msg: "6df067add738195fd55ac2e76b476971b9a0e6d8", Tag: "5cb595f9587afa7470a3157040b917bf", Result: "valid"},

{TcId: 20, Flag: "Pseudorandom", Key: "fda6a01194beb462953d7e6c49b32dac", Msg: "f60ae3b036abcab78c98fc1d4b67970c0955cb6fe24483f8907fd73319679b", Tag: "1f0f8124ab6c832e87684bac701544c1", Result: "valid"},

{TcId: 21, Flag: "Pseudorandom", Key: "9bd3902ed0996c869b572272e76f3889", Msg: "a7ba19d49ee1ea02f098aa8e30c740d893a4456ccc294040484ed8a00a55f93e", Tag: "45082218c2d05eef32247feb1133d0a3", Result: "valid"},

{TcId: 22, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "96dd6e5a882cbd564c39ae7d1c5a31aa", Result: "invalid"},

{TcId: 23, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "43802eb1931f0032afe984443738cd31", Result: "invalid"},

{TcId: 24, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7acfbbca7a2ea68b966fc5399f74809e", Result: "invalid"},

{TcId: 25, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "95dd6e5a882cbd564c39ae7d1c5a31aa", Result: "invalid"},

{TcId: 26, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "40802eb1931f0032afe984443738cd31", Result: "invalid"},

{TcId: 27, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "79cfbbca7a2ea68b966fc5399f74809e", Result: "invalid"},

{TcId: 28, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "17dd6e5a882cbd564c39ae7d1c5a31aa", Result: "invalid"},

{TcId: 29, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "c2802eb1931f0032afe984443738cd31", Result: "invalid"},

{TcId: 30, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "fbcfbbca7a2ea68b966fc5399f74809e", Result: "invalid"},

{TcId: 31, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dc6e5a882cbd564c39ae7d1c5a31aa", Result: "invalid"},

{TcId: 32, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42812eb1931f0032afe984443738cd31", Result: "invalid"},

{TcId: 33, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcebbca7a2ea68b966fc5399f74809e", Result: "invalid"},

{TcId: 34, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6eda882cbd564c39ae7d1c5a31aa", Result: "invalid"},

{TcId: 35, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802e31931f0032afe984443738cd31", Result: "invalid"},

{TcId: 36, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbb4a7a2ea68b966fc5399f74809e", Result: "invalid"},

{TcId: 37, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a892cbd564c39ae7d1c5a31aa", Result: "invalid"},

{TcId: 38, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1921f0032afe984443738cd31", Result: "invalid"},

{TcId: 39, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7b2ea68b966fc5399f74809e", Result: "invalid"},

{TcId: 40, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a8a2cbd564c39ae7d1c5a31aa", Result: "invalid"},

{TcId: 41, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1911f0032afe984443738cd31", Result: "invalid"},

{TcId: 42, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca782ea68b966fc5399f74809e", Result: "invalid"},

{TcId: 43, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbdd64c39ae7d1c5a31aa", Result: "invalid"},

{TcId: 44, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f00b2afe984443738cd31", Result: "invalid"},

{TcId: 45, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea60b966fc5399f74809e", Result: "invalid"},

{TcId: 46, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd564d39ae7d1c5a31aa", Result: "invalid"},

{TcId: 47, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f0032aee984443738cd31", Result: "invalid"},

{TcId: 48, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b976fc5399f74809e", Result: "invalid"},

{TcId: 49, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd56cc39ae7d1c5a31aa", Result: "invalid"},

{TcId: 50, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f00322fe984443738cd31", Result: "invalid"},

{TcId: 51, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b166fc5399f74809e", Result: "invalid"},

{TcId: 52, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd564c19ae7d1c5a31aa", Result: "invalid"},

{TcId: 53, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f0032afc984443738cd31", Result: "invalid"},

{TcId: 54, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b964fc5399f74809e", Result: "invalid"},

{TcId: 55, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd564c39af7d1c5a31aa", Result: "invalid"},

{TcId: 56, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f0032afe985443738cd31", Result: "invalid"},

{TcId: 57, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b966fc4399f74809e", Result: "invalid"},

{TcId: 58, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd564c39ae7d1d5a31aa", Result: "invalid"},

{TcId: 59, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f0032afe984443638cd31", Result: "invalid"},

{TcId: 60, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b966fc5399e74809e", Result: "invalid"},

{TcId: 61, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd564c39ae7d1e5a31aa", Result: "invalid"},

{TcId: 62, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f0032afe984443538cd31", Result: "invalid"},

{TcId: 63, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b966fc5399d74809e", Result: "invalid"},

{TcId: 64, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd564c39ae7d9c5a31aa", Result: "invalid"},

{TcId: 65, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f0032afe98444b738cd31", Result: "invalid"},

{TcId: 66, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b966fc5391f74809e", Result: "invalid"},

{TcId: 67, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd564c39ae7d1c5a31ab", Result: "invalid"},

{TcId: 68, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f0032afe984443738cd30", Result: "invalid"},

{TcId: 69, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b966fc5399f74809f", Result: "invalid"},

{TcId: 70, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd564c39ae7d1c5a31a8", Result: "invalid"},

{TcId: 71, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f0032afe984443738cd33", Result: "invalid"},

{TcId: 72, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b966fc5399f74809c", Result: "invalid"},

{TcId: 73, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd564c39ae7d1c5a31ea", Result: "invalid"},

{TcId: 74, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f0032afe984443738cd71", Result: "invalid"},

{TcId: 75, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b966fc5399f7480de", Result: "invalid"},

{TcId: 76, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbd564c39ae7d1c5a312a", Result: "invalid"},

{TcId: 77, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f0032afe984443738cdb1", Result: "invalid"},

{TcId: 78, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea68b966fc5399f74801e", Result: "invalid"},

{TcId: 79, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "96dd6e5a882cbd564d39ae7d1c5a31aa", Result: "invalid"},

{TcId: 80, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "43802eb1931f0032aee984443738cd31", Result: "invalid"},

{TcId: 81, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7acfbbca7a2ea68b976fc5399f74809e", Result: "invalid"},

{TcId: 82, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6eda882cbdd64c39ae7d1c5a31aa", Result: "invalid"},

{TcId: 83, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802e31931f00b2afe984443738cd31", Result: "invalid"},

{TcId: 84, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbb4a7a2ea60b966fc5399f74809e", Result: "invalid"},

{TcId: 85, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "97dd6e5a882cbdd64c39ae7d1c5a312a", Result: "invalid"},

{TcId: 86, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "42802eb1931f00b2afe984443738cdb1", Result: "invalid"},

{TcId: 87, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7bcfbbca7a2ea60b966fc5399f74801e", Result: "invalid"},

{TcId: 88, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "682291a577d342a9b3c65182e3a5ce55", Result: "invalid"},

{TcId: 89, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "bd7fd14e6ce0ffcd50167bbbc8c732ce", Result: "invalid"},

{TcId: 90, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "8430443585d1597469903ac6608b7f61", Result: "invalid"},

{TcId: 91, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "00000000000000000000000000000000", Result: "invalid"},

{TcId: 92, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "00000000000000000000000000000000", Result: "invalid"},

{TcId: 93, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "00000000000000000000000000000000", Result: "invalid"},

{TcId: 94, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "ffffffffffffffffffffffffffffffff", Result: "invalid"},

{TcId: 95, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "ffffffffffffffffffffffffffffffff", Result: "invalid"},

{TcId: 96, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "ffffffffffffffffffffffffffffffff", Result: "invalid"},

{TcId: 97, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "175deeda08ac3dd6ccb92efd9cdab12a", Result: "invalid"},

{TcId: 98, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "c200ae31139f80b22f6904c4b7b84db1", Result: "invalid"},

{TcId: 99, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "fb4f3b4afaae260b16ef45b91ff4001e", Result: "invalid"},

{TcId: 100, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "", Tag: "96dc6f5b892dbc574d38af7c1d5b30ab", Result: "invalid"},

{TcId: 101, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "0001020304050607", Tag: "43812fb0921e0133aee885453639cc30", Result: "invalid"},

{TcId: 102, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "7acebacb7b2fa78a976ec4389e75819f", Result: "invalid"},

{TcId: 103, Flag: "Pseudorandom", Key: "3d6bf9edae6d881eade0ff8c7076a4835b71320c1f36b631", Msg: "", Tag: "a8dd15fe2ce3495ec5b666744ec29220", Result: "valid"},

{TcId: 104, Flag: "Pseudorandom", Key: "915429743435c28997a33b33b6574a953d81dae0e7032e6a", Msg: "58", Tag: "e13b3f7f7f510c3a059df7a68c7e2ad5", Result: "valid"},

{TcId: 105, Flag: "Pseudorandom", Key: "f0c288ba26b284f9fb321b444a6517b3cdda1a799d55fdff", Msg: "0f7e", Tag: "06ef847f5f9dbf03a4f283da8c400220", Result: "valid"},

{TcId: 106, Flag: "Pseudorandom", Key: "6b55e4d4fd6847a80a6bfb0dcc0aa93f9fd797fc5c50292e", Msg: "33f530", Tag: "dd135053a47ca8f282c299e83b8c57c4", Result: "valid"},

{TcId: 107, Flag: "Pseudorandom", Key: "1eb21a9e995a8e45c9e71ecbd6fe615b3e0318007c64b644", Msg: "3aa73c48", Tag: "1e93fff846934a6eea0575eecb0f0e1f", Result: "valid"},

{TcId: 108, Flag: "Pseudorandom", Key: "710e2d5d4a9f0bc7e50796655e046a18cc5769d7764355da", Msg: "7e4c690a88", Tag: "016d4df06c68a6a788a9ea052e1b550d", Result: "valid"},

{TcId: 109, Flag: "Pseudorandom", Key: "d8c09ea400779b63e774bdacd0cb7b5dd6f736ca23d52acf", Msg: "e9520280973b", Tag: "8030ae9f98f5d20c6089f6b1bd87c29e", Result: "valid"},

{TcId: 110, Flag: "Pseudorandom", Key: "8e67e9a0863b55bed408866f1cbc05357abe3f9d79f406f2", Msg: "4880b412287a0b", Tag: "bcaf50785f062a8fb8dd3c2c4cead2e1", Result: "valid"},

{TcId: 111, Flag: "Pseudorandom", Key: "28d8da67806410e5565bcc5a9d7ab9fb357413fa0158378c", Msg: "004e3f4a4e6db955", Tag: "c4c2c0876be9eabeb5a956da53846b08", Result: "valid"},

{TcId: 112, Flag: "Pseudorandom", Key: "dc968dd89fd602bb7eca6f3a8a13e4f59c08d02a514b1934", Msg: "41a25354efeb1bc3b8", Tag: "f33a62caf397f9aff71fe42941ba41d8", Result: "valid"},

{TcId: 113, Flag: "Pseudorandom", Key: "7658951c0f620d82afd92756cc2d7983b79da3e56fdd1b78", Msg: "f0e82fb5c5666f4af49f", Tag: "4d724d05f3402967eb65ae1e32d5469e", Result: "valid"},

{TcId: 114, Flag: "Pseudorandom", Key: "d9574c3a221b986690931faac5258d9d3c52362b2cb9b054", Msg: "178ea8404ba54ee4e4522c", Tag: "64a0e0b6757309ab58d74f72c310e473", Result: "valid"},

{TcId: 115, Flag: "Pseudorandom", Key: "704409bab28085c44981f28f75dd143a4f747106f63f262e", Msg: "cda5709e7f115624e74ab031", Tag: "6ab2074334be14a95b6a241f897a43de", Result: "valid"},

{TcId: 116, Flag: "Pseudorandom", Key: "d8d06ef6a53bbff5c8f12d791b8f4c67e574bf440736d1cc", Msg: "a1171eae1979f48345dd9485a0", Tag: "7aa57cf98b24897cc9230e3316758e61", Result: "valid"},

{TcId: 117, Flag: "Pseudorandom", Key: "71129e781613f39d9ac39fbde2628b44c250c14deb5ef9e2", Msg: "967593cc64bcbf7f3c58d04cb82b", Tag: "6cc488b0a40eadbe4bcee2623239d126", Result: "valid"},

{TcId: 118, Flag: "Pseudorandom", Key: "850fc859e9f7b89a367611dee6698f33962d8245ca8dc331", Msg: "586f4f171af116519061a8e0e77940", Tag: "fb11a360c9776991d73d6e41d07710a2", Result: "valid"},

{TcId: 119, Flag: "Pseudorandom", Key: "f4bfa5aa4f0f4d62cf736cd2969c43d580fdb92f2753bedb", Msg: "0e239f239705b282ce2200fe20de1165", Tag: "ab20a6cf60873665b1d6999b05c7f9c6", Result: "valid"},

{TcId: 120, Flag: "Pseudorandom", Key: "cfd3f68873d81a27d2bfce876c79f6e609074dec39e34614", Msg: "b1973cb25aa87ef9d1a8888b0a0f5c04c6", Tag: "b95a016b83a0ae4194023333c8a7345a", Result: "valid"},

{TcId: 121, Flag: "Pseudorandom", Key: "b7f165bced1613da5e747fdf9255832d30c07f2deeb5a326", Msg: "289647ea8d0ff31375a82aa1c620903048bb1d0e", Tag: "3b1e84eb3d4a2233caf1982905940393", Result: "valid"},

{TcId: 122, Flag: "Pseudorandom", Key: "9bbe6e004fb260dadb02b68b78954f1da5e6a2d02e0aeefe", Msg: "665423092ce95b927e98b8082030f58e33f3ec1b0c29532c2f421855f00f97", Tag: "0e434cfb3d0ef0584e03bd5648934df6", Result: "valid"},

{TcId: 123, Flag: "Pseudorandom", Key: "9d11abc1fcb248a436598e695be12c3c2ed90a18ba09d62c", Msg: "aa5182cae2a8fb068c0b3fb2be3e57ae523d13dffd1a944587707c2b67447f3f", Tag: "8597d9a04d1c271d61d42f007b435175", Result: "valid"},

{TcId: 124, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ed12390ea0a7ed15d9d37a6eca1fc990", Result: "invalid"},

{TcId: 125, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c81307df60859acb911c7be61be7ca90", Result: "invalid"},

{TcId: 126, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f91bde0069a6e389573bf04e7cde688c", Result: "invalid"},

{TcId: 127, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ee12390ea0a7ed15d9d37a6eca1fc990", Result: "invalid"},

{TcId: 128, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "cb1307df60859acb911c7be61be7ca90", Result: "invalid"},

{TcId: 129, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "fa1bde0069a6e389573bf04e7cde688c", Result: "invalid"},

{TcId: 130, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "6c12390ea0a7ed15d9d37a6eca1fc990", Result: "invalid"},

{TcId: 131, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "491307df60859acb911c7be61be7ca90", Result: "invalid"},

{TcId: 132, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "781bde0069a6e389573bf04e7cde688c", Result: "invalid"},

{TcId: 133, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec13390ea0a7ed15d9d37a6eca1fc990", Result: "invalid"},

{TcId: 134, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91207df60859acb911c7be61be7ca90", Result: "invalid"},

{TcId: 135, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81ade0069a6e389573bf04e7cde688c", Result: "invalid"},

{TcId: 136, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12398ea0a7ed15d9d37a6eca1fc990", Result: "invalid"},

{TcId: 137, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c913075f60859acb911c7be61be7ca90", Result: "invalid"},

{TcId: 138, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde8069a6e389573bf04e7cde688c", Result: "invalid"},

{TcId: 139, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea1a7ed15d9d37a6eca1fc990", Result: "invalid"},

{TcId: 140, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df61859acb911c7be61be7ca90", Result: "invalid"},

{TcId: 141, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0068a6e389573bf04e7cde688c", Result: "invalid"},

{TcId: 142, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea2a7ed15d9d37a6eca1fc990", Result: "invalid"},

{TcId: 143, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df62859acb911c7be61be7ca90", Result: "invalid"},

{TcId: 144, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde006ba6e389573bf04e7cde688c", Result: "invalid"},

{TcId: 145, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed95d9d37a6eca1fc990", Result: "invalid"},

{TcId: 146, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859a4b911c7be61be7ca90", Result: "invalid"},

{TcId: 147, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e309573bf04e7cde688c", Result: "invalid"},

{TcId: 148, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed15d8d37a6eca1fc990", Result: "invalid"},

{TcId: 149, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb901c7be61be7ca90", Result: "invalid"},

{TcId: 150, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389563bf04e7cde688c", Result: "invalid"},

{TcId: 151, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed1559d37a6eca1fc990", Result: "invalid"},

{TcId: 152, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb111c7be61be7ca90", Result: "invalid"},

{TcId: 153, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389d73bf04e7cde688c", Result: "invalid"},

{TcId: 154, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed15d9f37a6eca1fc990", Result: "invalid"},

{TcId: 155, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb913c7be61be7ca90", Result: "invalid"},

{TcId: 156, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389571bf04e7cde688c", Result: "invalid"},

{TcId: 157, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed15d9d37b6eca1fc990", Result: "invalid"},

{TcId: 158, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb911c7ae61be7ca90", Result: "invalid"},

{TcId: 159, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389573bf14e7cde688c", Result: "invalid"},

{TcId: 160, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed15d9d37a6ecb1fc990", Result: "invalid"},

{TcId: 161, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb911c7be61ae7ca90", Result: "invalid"},

{TcId: 162, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389573bf04e7dde688c", Result: "invalid"},

{TcId: 163, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed15d9d37a6ec81fc990", Result: "invalid"},

{TcId: 164, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb911c7be619e7ca90", Result: "invalid"},

{TcId: 165, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389573bf04e7ede688c", Result: "invalid"},

{TcId: 166, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed15d9d37a6e4a1fc990", Result: "invalid"},

{TcId: 167, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb911c7be69be7ca90", Result: "invalid"},

{TcId: 168, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389573bf04efcde688c", Result: "invalid"},

{TcId: 169, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed15d9d37a6eca1fc991", Result: "invalid"},

{TcId: 170, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb911c7be61be7ca91", Result: "invalid"},

{TcId: 171, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389573bf04e7cde688d", Result: "invalid"},

{TcId: 172, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed15d9d37a6eca1fc992", Result: "invalid"},

{TcId: 173, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb911c7be61be7ca92", Result: "invalid"},

{TcId: 174, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389573bf04e7cde688e", Result: "invalid"},

{TcId: 175, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed15d9d37a6eca1fc9d0", Result: "invalid"},

{TcId: 176, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb911c7be61be7cad0", Result: "invalid"},

{TcId: 177, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389573bf04e7cde68cc", Result: "invalid"},

{TcId: 178, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed15d9d37a6eca1fc910", Result: "invalid"},

{TcId: 179, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859acb911c7be61be7ca10", Result: "invalid"},

{TcId: 180, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e389573bf04e7cde680c", Result: "invalid"},

{TcId: 181, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ed12390ea0a7ed15d8d37a6eca1fc990", Result: "invalid"},

{TcId: 182, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c81307df60859acb901c7be61be7ca90", Result: "invalid"},

{TcId: 183, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f91bde0069a6e389563bf04e7cde688c", Result: "invalid"},

{TcId: 184, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12398ea0a7ed95d9d37a6eca1fc990", Result: "invalid"},

{TcId: 185, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c913075f60859a4b911c7be61be7ca90", Result: "invalid"},

{TcId: 186, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde8069a6e309573bf04e7cde688c", Result: "invalid"},

{TcId: 187, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ec12390ea0a7ed95d9d37a6eca1fc910", Result: "invalid"},

{TcId: 188, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c91307df60859a4b911c7be61be7ca10", Result: "invalid"},

{TcId: 189, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f81bde0069a6e309573bf04e7cde680c", Result: "invalid"},

{TcId: 190, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "13edc6f15f5812ea262c859135e0366f", Result: "invalid"},

{TcId: 191, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "36ecf8209f7a65346ee38419e418356f", Result: "invalid"},

{TcId: 192, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "07e421ff96591c76a8c40fb183219773", Result: "invalid"},

{TcId: 193, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "00000000000000000000000000000000", Result: "invalid"},

{TcId: 194, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "00000000000000000000000000000000", Result: "invalid"},

{TcId: 195, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "00000000000000000000000000000000", Result: "invalid"},

{TcId: 196, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ffffffffffffffffffffffffffffffff", Result: "invalid"},

{TcId: 197, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "ffffffffffffffffffffffffffffffff", Result: "invalid"},

{TcId: 198, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "ffffffffffffffffffffffffffffffff", Result: "invalid"},

{TcId: 199, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "6c92b98e20276d955953faee4a9f4910", Result: "invalid"},

{TcId: 200, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "4993875fe0051a4b119cfb669b674a10", Result: "invalid"},

{TcId: 201, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "789b5e80e9266309d7bb70cefc5ee80c", Result: "invalid"},

{TcId: 202, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "", Tag: "ed13380fa1a6ec14d8d27b6fcb1ec891", Result: "invalid"},

{TcId: 203, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "0001020304050607", Tag: "c81206de61849bca901d7ae71ae6cb91", Result: "invalid"},

{TcId: 204, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f1011121314151617", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "f91adf0168a7e288563af14f7ddf698d", Result: "invalid"},

{TcId: 205, Flag: "Pseudorandom", Key: "7bf9e536b66a215c22233fe2daaa743a898b9acb9f7802de70b40e3d6e43ef97", Msg: "", Tag: "736c7b56957db774c5ddf7c7a70ba8a8", Result: "valid"},

{TcId: 206, Flag: "Pseudorandom", Key: "e754076ceab3fdaf4f9bcab7d4f0df0cbbafbc87731b8f9b7cd2166472e8eebc", Msg: "40", Tag: "9d47482c2d9252bace43a75a8335b8b8", Result: "valid"},

{TcId: 207, Flag: "Pseudorandom", Key: "ea3b016bdd387dd64d837c71683808f335dbdc53598a4ea8c5f952473fafaf5f", Msg: "6601", Tag: "c7c44e31c466334992d6f9de3c771634", Result: "valid"},

{TcId: 208, Flag: "Pseudorandom", Key: "73d4709637857dafab6ad8b2b0a51b06524717fedf100296644f7cfdaae1805b", Msg: "f1d300", Tag: "b7086603a85e11fceb8cadea9bd30939", Result: "valid"},

{TcId: 209, Flag: "Pseudorandom", Key: "d5c81b399d4c0d1583a13da56de6d2dc45a66e7b47c24ab1192e246dc961dd77", Msg: "2ae63cbf", Tag: "ba383a3a15c9df64bba50d611113a024", Result: "valid"},

{TcId: 210, Flag: "Pseudorandom", Key: "2521203fa0dddf59d837b2830f87b1aa61f958155df3ca4d1df2457cb4284dc8", Msg: "af3a015ea1", Tag: "b457137c548908c629f714fe83b1ed90", Result: "valid"},

{TcId: 211, Flag: "Pseudorandom", Key: "665a02bc265a66d01775091da56726b6668bfd903cb7af66fb1b78a8a062e43c", Msg: "3f56935def3f", Tag: "b6d6fde93fc85de289b36b446d77b423", Result: "valid"},

{TcId: 212, Flag: "Pseudorandom", Key: "facd75b22221380047305bc981f570e2a1af38928ea7e2059e3af5fc6b82b493", Msg: "57bb86beed156f", Tag: "8b1ef72d0a612735b08efef981f213c2", Result: "valid"},

{TcId: 213, Flag: "Pseudorandom", Key: "505aa98819809ef63b9a368a1e8bc2e922da45b03ce02d9a7966b15006dba2d5", Msg: "2e4e7ef728fe11af", Tag: "f79606b83a7706a2a19e068bce818898", Result: "valid"},

{TcId: 214, Flag: "Pseudorandom", Key: "f942093842808ba47f64e427f7351dde6b9546e66de4e7d60aa6f328182712cf", Msg: "852a21d92848e627c7", Tag: "a5a877f22ac743b7fb9e050d2e3ddb02", Result: "valid"},

{TcId: 215, Flag: "Pseudorandom", Key: "64be162b39c6e5f1fed9c32d9f674d9a8cde6eaa2443214d86bd4a1fb53b81b4", Msg: "195a3b292f93baff0a2c", Tag: "6ea172e5c4d2fac075ca602de5757a62", Result: "valid"},

{TcId: 216, Flag: "Pseudorandom", Key: "b259a555d44b8a20c5489e2f38392ddaa6be9e35b9833b67e1b5fdf6cb3e4c6c", Msg: "afd73117330c6e8528a6e4", Tag: "68020bfc9bd73fd80d3ce581ba3b1208", Result: "valid"},

{TcId: 217, Flag: "Pseudorandom", Key: "2c6fc62daa77ba8c6881b3dd6989898fef646663cc7b0a3db8228a707b85f2dc", Msg: "0ff54d6b6759120c2e8a51e3", Tag: "110edd727a9bf7fa11a6358afe617d9d", Result: "valid"},

{TcId: 218, Flag: "Pseudorandom", Key: "abab815d51df29f740e4e2079fb798e0152836e6ab57d1536ae8929e52c06eb8", Msg: "f0058d412a104e53d820b95a7f", Tag: "1fa24c6625a0f8e1fc37827ac84d3cc4", Result: "valid"},

{TcId: 219, Flag: "Pseudorandom", Key: "3d5da1af83f7287458bff7a7651ea5d8db72259401333f6b82096996dd7eaf19", Msg: "aacc36972f183057919ff57b49e1", Tag: "868765a8fa6aa898ddec0f4123e996be", Result: "valid"},

{TcId: 220, Flag: "Pseudorandom", Key: "c19bdf314c6cf64381425467f42aefa17c1cc9358be16ce31b1d214859ce86aa", Msg: "5d066a92c300e9b6ddd63a7c13ae33", Tag: "b96818b7acaf879c7a7f8271375a6914", Result: "valid"},

{TcId: 221, Flag: "Pseudorandom", Key: "612e837843ceae7f61d49625faa7e7494f9253e20cb3adcea686512b043936cd", Msg: "cc37fae15f745a2f40e2c8b192f2b38d", Tag: "4b88e193000c5a4b23e95c7f2b26530b", Result: "valid"},

{TcId: 222, Flag: "Pseudorandom", Key: "73216fafd0022d0d6ee27198b2272578fa8f04dd9f44467fbb6437aa45641bf7", Msg: "d5247b8f6c3edcbfb1d591d13ece23d2f5", Tag: "86911c7da51dc0823d6e93d4290d1ad4", Result: "valid"},

{TcId: 223, Flag: "Pseudorandom", Key: "c2039f0d05951aa8d9fbdf68be58a37cf99bd1afcedda286a9db470c3729ca92", Msg: "ed5b5e28e9703bdf5c7b3b080f2690a605fcd0d9", Tag: "24e1f4416b9980ef4c2795e9c4bf503f", Result: "valid"},

{TcId: 224, Flag: "Pseudorandom", Key: "4f097858a1aec62cf18f0966b2b120783aa4ae9149d3213109740506ae47adfe", Msg: "ee53d8e5039e82d9fcca114e375a014febfea117a7e709d9008d43858e3660", Tag: "a5a66fa3aa3dabe032d77f438457c056", Result: "valid"},

{TcId: 225, Flag: "Pseudorandom", Key: "96e1e4896fb2cd05f133a6a100bc5609a7ac3ca6d81721e922dadd69ad07a892", Msg: "91a17e4dfcc3166a1add26ff0e7c12056e8a654f28a6de24f4ba739ceb5b5b18", Tag: "925f177d85ea297ef14b203fe409f9ab", Result: "valid"},

{TcId: 226, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6af0a293d8cba0101f0089727691b7fb", Result: "invalid"},

{TcId: 227, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d709717c3a4ef8a2ea200b297d2accec", Result: "invalid"},

{TcId: 228, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "58ee3f3b5f83e290cae26dad29bba32d", Result: "invalid"},

{TcId: 229, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "69f0a293d8cba0101f0089727691b7fb", Result: "invalid"},

{TcId: 230, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d409717c3a4ef8a2ea200b297d2accec", Result: "invalid"},

{TcId: 231, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "5bee3f3b5f83e290cae26dad29bba32d", Result: "invalid"},

{TcId: 232, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "ebf0a293d8cba0101f0089727691b7fb", Result: "invalid"},

{TcId: 233, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "5609717c3a4ef8a2ea200b297d2accec", Result: "invalid"},

{TcId: 234, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "d9ee3f3b5f83e290cae26dad29bba32d", Result: "invalid"},

{TcId: 235, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf1a293d8cba0101f0089727691b7fb", Result: "invalid"},

{TcId: 236, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d608717c3a4ef8a2ea200b297d2accec", Result: "invalid"},

{TcId: 237, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ef3f3b5f83e290cae26dad29bba32d", Result: "invalid"},

{TcId: 238, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a213d8cba0101f0089727691b7fb", Result: "invalid"},

{TcId: 239, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d60971fc3a4ef8a2ea200b297d2accec", Result: "invalid"},

{TcId: 240, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3fbb5f83e290cae26dad29bba32d", Result: "invalid"},

{TcId: 241, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d9cba0101f0089727691b7fb", Result: "invalid"},

{TcId: 242, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3b4ef8a2ea200b297d2accec", Result: "invalid"},

{TcId: 243, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5e83e290cae26dad29bba32d", Result: "invalid"},

{TcId: 244, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293dacba0101f0089727691b7fb", Result: "invalid"},

{TcId: 245, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c384ef8a2ea200b297d2accec", Result: "invalid"},

{TcId: 246, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5d83e290cae26dad29bba32d", Result: "invalid"},

{TcId: 247, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0901f0089727691b7fb", Result: "invalid"},

{TcId: 248, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef822ea200b297d2accec", Result: "invalid"},

{TcId: 249, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e210cae26dad29bba32d", Result: "invalid"},

{TcId: 250, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0101e0089727691b7fb", Result: "invalid"},

{TcId: 251, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a2eb200b297d2accec", Result: "invalid"},

{TcId: 252, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e290cbe26dad29bba32d", Result: "invalid"},

{TcId: 253, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0109f0089727691b7fb", Result: "invalid"},

{TcId: 254, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a26a200b297d2accec", Result: "invalid"},

{TcId: 255, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e2904ae26dad29bba32d", Result: "invalid"},

{TcId: 256, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0101f2089727691b7fb", Result: "invalid"},

{TcId: 257, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a2ea000b297d2accec", Result: "invalid"},

{TcId: 258, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e290cac26dad29bba32d", Result: "invalid"},

{TcId: 259, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0101f0088727691b7fb", Result: "invalid"},

{TcId: 260, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a2ea200a297d2accec", Result: "invalid"},

{TcId: 261, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e290cae26cad29bba32d", Result: "invalid"},

{TcId: 262, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0101f0089727791b7fb", Result: "invalid"},

{TcId: 263, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a2ea200b297c2accec", Result: "invalid"},

{TcId: 264, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e290cae26dad28bba32d", Result: "invalid"},

{TcId: 265, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0101f0089727491b7fb", Result: "invalid"},

{TcId: 266, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a2ea200b297f2accec", Result: "invalid"},

{TcId: 267, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e290cae26dad2bbba32d", Result: "invalid"},

{TcId: 268, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0101f008972f691b7fb", Result: "invalid"},

{TcId: 269, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a2ea200b29fd2accec", Result: "invalid"},

{TcId: 270, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e290cae26dada9bba32d", Result: "invalid"},

{TcId: 271, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0101f0089727691b7fa", Result: "invalid"},

{TcId: 272, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a2ea200b297d2acced", Result: "invalid"},

{TcId: 273, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e290cae26dad29bba32c", Result: "invalid"},

{TcId: 274, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0101f0089727691b7f9", Result: "invalid"},

{TcId: 275, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a2ea200b297d2accee", Result: "invalid"},

{TcId: 276, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e290cae26dad29bba32f", Result: "invalid"},

{TcId: 277, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0101f0089727691b7bb", Result: "invalid"},

{TcId: 278, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a2ea200b297d2accac", Result: "invalid"},

{TcId: 279, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e290cae26dad29bba36d", Result: "invalid"},

{TcId: 280, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0101f0089727691b77b", Result: "invalid"},

{TcId: 281, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef8a2ea200b297d2acc6c", Result: "invalid"},

{TcId: 282, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e290cae26dad29bba3ad", Result: "invalid"},

{TcId: 283, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6af0a293d8cba0101e0089727691b7fb", Result: "invalid"},

{TcId: 284, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d709717c3a4ef8a2eb200b297d2accec", Result: "invalid"},

{TcId: 285, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "58ee3f3b5f83e290cbe26dad29bba32d", Result: "invalid"},

{TcId: 286, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a213d8cba0901f0089727691b7fb", Result: "invalid"},

{TcId: 287, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d60971fc3a4ef822ea200b297d2accec", Result: "invalid"},

{TcId: 288, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3fbb5f83e210cae26dad29bba32d", Result: "invalid"},

{TcId: 289, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6bf0a293d8cba0901f0089727691b77b", Result: "invalid"},

{TcId: 290, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d609717c3a4ef822ea200b297d2acc6c", Result: "invalid"},

{TcId: 291, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "59ee3f3b5f83e210cae26dad29bba3ad", Result: "invalid"},

{TcId: 292, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "940f5d6c27345fefe0ff768d896e4804", Result: "invalid"},

{TcId: 293, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "29f68e83c5b1075d15dff4d682d53313", Result: "invalid"},

{TcId: 294, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "a611c0c4a07c1d6f351d9252d6445cd2", Result: "invalid"},

{TcId: 295, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "00000000000000000000000000000000", Result: "invalid"},

{TcId: 296, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "00000000000000000000000000000000", Result: "invalid"},

{TcId: 297, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "00000000000000000000000000000000", Result: "invalid"},

{TcId: 298, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "ffffffffffffffffffffffffffffffff", Result: "invalid"},

{TcId: 299, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "ffffffffffffffffffffffffffffffff", Result: "invalid"},

{TcId: 300, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "ffffffffffffffffffffffffffffffff", Result: "invalid"},

{TcId: 301, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "eb702213584b20909f8009f2f611377b", Result: "invalid"},

{TcId: 302, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "5689f1fcbace78226aa08ba9fdaa4c6c", Result: "invalid"},

{TcId: 303, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "d96ebfbbdf0362104a62ed2da93b23ad", Result: "invalid"},

{TcId: 304, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "", Tag: "6af1a392d9caa1111e0188737790b6fa", Result: "invalid"},

{TcId: 305, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "0001020304050607", Tag: "d708707d3b4ff9a3eb210a287c2bcded", Result: "invalid"},

{TcId: 306, Flag: "ModifiedTag", Key: "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", Msg: "000102030405060708090a0b0c0d0e0f", Tag: "58ef3e3a5e82e391cbe36cac28baa22c", Result: "invalid"},

{TcId: 307, Flag: "InvalidKeySize", Key: "", Msg: "00b9449326d39416", Tag: "", Result: "invalid"},

{TcId: 308, Flag: "InvalidKeySize", Key: "0f", Msg: "4538b79a1397e2aa", Tag: "", Result: "invalid"},

{TcId: 309, Flag: "InvalidKeySize", Key: "a88e385af7185148", Msg: "dc63b7ef08096e4f", Tag: "", Result: "invalid"},

{TcId: 310, Flag: "InvalidKeySize", Key: "003a228008d390b645929df73a2b2bdd8298918d", Msg: "ad1d3c3122ab7ac6", Tag: "", Result: "invalid"},

{TcId: 311, Flag: "InvalidKeySize", Key: "94baaac150e2645ae1ec1939c7bcefb73f6edb146fae02289b6c6326ff39bc265d612bef2727fa72", Msg: "e3f75a886c4a5591", Tag: "", Result: "invalid"},

}
