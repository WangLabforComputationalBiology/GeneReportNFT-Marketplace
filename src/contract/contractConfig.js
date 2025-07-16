/**合约配置
 * @param {json} abi 合约ABI，对于智能合约的描述。格式或为.abi，目前直接修改后缀即可
 * @param {string} contractAddress 合约地址，不同的合约目前需要手动更新
 */
import abi from "@/contract/contractABI.json"
const contractAddress = '0xB63F003d7464F840F7992fAD36e57900DB1A40c0';
const contractABI = abi;

export { contractAddress, contractABI }