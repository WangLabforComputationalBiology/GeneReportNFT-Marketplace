import { useWalletStore } from "@/stores/account";

/* 无登录状态拦截 */
const walletAuthGuard = (to, from, next) => {
    const walletAddress = useWalletStore();
    if (!walletAddress.address && to.meta.requiresAuth) {
        next("/login"); // 重定向到登录页
        alert("No account accessible.Please setup your wallet first.");
    } else {
        next(); // 放行
    }
};

export default walletAuthGuard;