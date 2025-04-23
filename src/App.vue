<template>
  <!-- 头部导航栏 -->
  <header>
    <div class="side" @click="toHome">
      LOGO
    </div>
    <div class="headerWrapper">
      <!-- 路由菜单 -->
      <div class="routers">
        <router-link to="/market" class="routerSelection ">Market</router-link>
        <router-link to="/create" class="routerSelection">Create</router-link>
        <!-- <router-link to="/drop" class="routerSelection">Drop</router-link> -->
        <router-link to="/stats" class="routerSelection">Stats</router-link>
      </div>

      <!-- 搜索框 -->
      <span class="navigation">
        <input id="navigationInput" type="text" placeholder="search..." />
        <div type="primary" class="searchBtn">
          <el-icon class="searchIcon" size="large">
            <Search />
          </el-icon>
        </div>
      </span>
      <button id="wegeneVerify"@click="redirectToOAuth">Wegene认证</button>
      <!-- 登录路由 -->
      <router-link to="/login" class="routerSelection" style="position:absolute;right:20px;width:50px">Login</router-link>
      <!-- 账户面板 -->
      <el-icon class="userIcon" v-if="account">
        <router-link to="/user" style="color: #169608;">
          <UserFilled />
        </router-link>
      </el-icon>
    </div>
    <span class="side">
    </span>
  </header>
  <!-- 路由主体 -->
  <router-view />
  <!-- 尾部导航栏 -->
  <footer>
    <!-- 内容主体 -->
    <div class="inBottom">
      <div class="communication">
        <p>Need some help? Please contact us.</p><br>
        <p>WangLab @SZTU</p>
        <p>Email: 11111@gmail.com</p>
        <p>Adress: xxxxxx</p>
      </div>
    </div>

  </footer>
</template>

<script>
import { useWalletStore } from '@/stores/account';


export default {
  name: "App",
  data() {
    return {
      account: null,
      provider: null,
      isVisible: false,
      accountDrawerVisible: false,
    };
  },

  created() {
    const wallet = useWalletStore()
    this.account = wallet.address;
    if (typeof window.ethereum !== 'undefined') {
      console.log('MetaMask is ready!');
    }

  },

  computed: {
    accountdisplay() {
      if (this.account.length > 10) {
        return this.account.slice(0, 10) + '...';
      }
      return this.account;
    },
  },

  methods: {
    //刷新返回主页
    toHome() {
      this.$router.push("/index");
    },
    Visible() {
      this.isVisible = true;
    },
    redirectToOAuth() {
      window.location.href = import.meta.env.VITE_APP_BASE_URL+'/user/oauth2Wegene';
    }
    
  },
}

</script>


<style lang="scss" scoped>
a {
  text-decoration: none;
}

header {
  z-index: 1000;
  width: 100%;
  height: 80px;
  display: flex;
  position: sticky;
  top: 0;
  transition: background-color 0.5s;
  background-color: #ffffffee;
  border-bottom: 1px solid #E4E7ED;

  .side {
    flex: 1;
    min-width: 120px;
    display: flex;
    align-items: center;
    cursor: pointer;
    font-size: 40px;
    color: #16952d;
    font-weight: bold;
  }

  .headerWrapper {
    align-items: center;
    position: relative;
    width: 80vw;
    min-width: 1200px;
    display: flex;
    margin: 0 auto;

    .navigation {
      width: 300px;
      display: flex;
      align-items: center;
      margin: 8px 0 0 20px;

      input {
        width: 250px;
        height: 35px;
        border: 2px solid #E4E7ED;
        border-radius: 10px;
      }

      .searchBtn {
        background-color: #169608;
        border-radius: 12px;
        margin-left: 4px;

        .searchIcon {
          height: 34px;
          width: 45px;
          color: #fff;
        }

        &:hover {
          background-color: #67C23A;
          color: #FFF;
          cursor: pointer;
        }
      }
    }

    .routers {
      height: 80px;
      display: flex;
    }

    .routerSelection {
      color: #169608;
      margin: 0 25px;
      font-size: 22px;
      align-content: center;
      font-weight: bold;

      &:hover {
        color: #67C23A;
      }

      &:first-child {
        margin-left: 0;
      }

    }


  }
}

.userIcon {
  position: absolute;
  right: 0;
  font-size: 30px;
  color: #169608;
  cursor: pointer;
}

//底部区域
footer {
  //绿色大背景
  width: 100%;
  height: 300px;
  background-color: #169608;
  margin: 0 auto;
  display: flex;

  //底部中心内容
  .inBottom {
    width: 1400px;
    margin: 50px auto;
    color: #FFF;
    font-size: 18px;

    .communication {
      width: 350px;
      height: 200px;
    }
  }
}
</style>
