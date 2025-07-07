/**合约配置
 * @param {json} abi 合约ABI，部署后由solidity IDEA返回，格式或为.abi，直接修改后缀即可
 */
import abi from "@/assets/sharingPlatform.json"
const contractAddress = '0xB63F003d7464F840F7992fAD36e57900DB1A40c0';
const contractABI = abi;

export { contractAddress, contractABI }