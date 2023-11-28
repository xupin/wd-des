package main

import (
	"encoding/hex"
	"fmt"

	"github.com/xupin/wd-des/utils"
)

const (
	desKey string = `548711fd`
	salt   string = `leiting`
)

func main() {
	hexStr := `662a14bd16bef3c6d5ab3e907fafed5c79d92a084c8ff5c9d2e1077601f2f7708b11d33ca7f494c60d6b0272a68dc4d1ce6dc4f35dc2e7909569da17d14ca3b81cbd2e4fa950d2ebc1fb1bb3146e23f5d5ab3e907fafed5c79d92a084c8ff5c9d2e1077601f2f770b64c50ba9d2425eff40a623f7712ccd3f0ca873e116d2c1bc32935b9e369e22aa17161018c3474321240e41187b7d7ebad46e44adc2987ecf2c7c38ba5d82be2587b1a7fcf6fdf3327f27a36ecc00bf129fc2096a7fae143ad46e44adc2987ecf6f70be24fd00aed04f926b6e409def154da85503182ab5c4827a053ab834371ecb66c8f6271bc5341df3e595ef3055c0ab105d3fb946010e53880e2cf2f69f55b6ca6042ea19d26b527cf69e333462ace41ceab8b526a0402d593efa1c72f091ba07257a7886e0cfb4ddec2c123c60d9ece97da872af219083162771fff1fe119c003682442d52bfb109b87419caa4d23c5354d58bb4f035f858e7d1e8c3f134802cd58212f7dec9134c722f85ad84a120854ff2317d17aab9dbd5de36ff29eeeddc6f31555f7d7ce41ceab8b526a0402d593efa1c72f099232b3fe87b7b1118034d6c8402d5a595d90cd35832bfdc8efab02e81b32d54f43d5021d2ebdcdbfe6ec4e7f49c360c8396f4533233c0b62a505f7fe557ad18314cb78557c72254603d70f7636b26b49183d798933b6993622810a007cc979d25be3a5af9de4b96e1044a21099fb6a70e0971dba3504c1420da388d82c3c5a9319c003682442d52b8f13e776fd5ec6a5b01483390221c8f2061b02f6d4886688e58fb53c7b60d1e099374a17fbd87ffea89f46faf00e0565f3fc60ff394f08ae8e43824d3348fab69d4ba7d57ced1357b5386899fdae6f727a4709e8897aee4f29f8800e7c30e60d6b44e67c20db51d12819ee15e8437043758b0ae30de3f10e49bb18f10404d97406b3961b7609a839895bd4bfc0baf909ea2c4da84430c50531dbe1826f585421a296c7aa7e3d445b8573b93db19fcde362d1c9f53dc7720ca4b4eb5a34eeca417a4709e8897aee4f32eaaa97fd635776e60116f080ea6c9b78cfd77220e56c2f808920f1151e9ffdac6599cfedb5ffaae4440d9b1a8717b64cd576b1b23952ccfc3be8f4b3e491b8ec9ee0b428dafba5fe8ed9998931693e1a5866fe37a12e4c8f9f858cf7f47c23885438079567dbb6a2e2183ab5e62cff1e06282cb6a360aba737635416adaa098711d73bd285ef11eb225555f5a07d324692547a40707861a0bc8a643e68616cdb6006e9278a76f6b6661f981dc80c60fd421c424a1528e352622bcf70d281007a4709e8897aee4fdb12465d875ade04bb5cc17531533fc97e8593f7220a6dfb36ae1a030ea20dc52a827afab9ac22bc9b12c8985e143ca61858d59cdf49aa1e604e7718ca72880986b5576dca424d0b9d189e40ce1214ffad44146c33eb91519a93f787c7924dbfb8c17f40c3ee6361f9de8957f44adfb43b92350ec577f06a4ca5f9223b230e858041c866b70bc2cc611684efa00161655e4ef90d64f2eed3a2c7fffc99ee06340055e6c25c162ee1ab25051ffec930d9077363a19507cc02ca8fd7e3754a9e316a7403b7eaf04d2ac018b63f2e3190a2e7aa448150ca3c3b0e43cb9052690e352b5d53fd8e68945c8e43824d3348fab6f54c2c0bf6c4b50d2ad7f3c0fc9081dd5fe78474f95d3c7a244be0a14d1078311e06282cb6a360ab0b1bc59cbd1eeb0e4f8c61251d119eb1523354e2c938e3a71d7014b5f729ca6e77532faa8566aa2fa175bb07192456d0b30688cd075bea42687214c02321cb5e5cb8572d6857d165876e73393a5b84b2bb0a710299a5148c73d87c79d99fe06532bfe030fb8586ee961e4894c6e5ed712e9af7fbf9e5543a`

	key := []byte(desKey)

	s := `{"loginUrl":"http://127.0.0.1:81/wd","payUrl":"http://127.0.0.1:81/wd","accountUrl":"http://127.0.0.1:81/wd","bbsUrl":"http://bbs.leiting.com","helpUrl":"http://helper.leiting.com","kernelUrl":"http://www.leiting.com","logUrl":"http://tplog.leiting.com","urlApis":[{"name":"login","value":"/login/login.php"},{"name":"checkLogin","value":"/login/mobile!mobileCheckLoginV3.php"},{"name":"fastRegister","value":"/login/mobile!fastRegisterV2.action"},{"name":"payOrder","value":"/terrace/phone_charge!createLeitingNo.action"},{"name":"getAdultInfo","value":"/terrace/game_api/getIdCardBindInfo.htm"}],"channels":[{"name":"leiting","value":"{\"guestLogin\":\"0\",\"payLevel\":\"1\",\"gmPhoneNum\":\"0592-3011618\"}"}],"plugs":[{"name":"qiyukf","value":"{\"appKey\":\"b8c33a8308b536dc4fde7be133fc9835\",\"groupId\":\"354781\",\"robotId\":\"83373\"}"},{"name":"toutiao","value":"{\"appId\":\"151304\",\"appName\":\"wendao\"}"},{"name":"wechath5","value":"{\"payResult\":\"/terrace/phone_charge!searchResultByLeitingNo.action\",\"payResultChannel\":\"/terrace/notify_back!searchResultByLeitingNo.action\",\"payLoginResult\":\"/terrace/drm_charge!queryWechatPayResult.action\"}"}],"resValues":[{"name":"lt_realname_auth_toast_msg1","value":"根据国家法规要求，您的账号信息尚未完善，请尽快完成实名认证以保障账号安全。"}]}`

	src, _ := hex.DecodeString(hexStr)
	bytes, _ := utils.DesDecrypt(src, key)

	fmt.Printf("解密内容: \n %s \n", string(bytes))

	bytes, _ = utils.DesEncrypt([]byte(s), key)

	hexStr = hex.EncodeToString(bytes)
	fmt.Printf("加密内容: \n %s \n", hexStr)

	fmt.Printf("加密hash: \n %s \n", utils.MD5(hexStr+salt))
}
