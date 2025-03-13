<template>
  <!-- 头部导航栏 -->
  <header>
    <span class="logo" @click="toHome">
      <p>LOGO</p>
    </span>

    <!-- 路由菜单 -->
    <div class="menu">
      <router-link to="/market" class="menuSelection">Market</router-link>
      <router-link to="/create" class="menuSelection">Create</router-link>
      <router-link to="/drop" class="menuSelection">Drop</router-link>
      <router-link to="/stats" class="menuSelection">Stats</router-link>
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

    <!-- 登录路由 -->
    <router-link to="/login" class="Login">Login</router-link>

  </header>
  <!-- 路由主体 -->
  <router-view />

  <!-- 尾部导航栏 -->
  <div class="bottom">
    <!-- 内容主体 -->
    <div class="inBottom">
      <div class="communication">
        <p>Need some help? Please contact us.</p><br>
        <p>WangLab @SZTU</p>
        <p>Email: 11111@gmail.com</p>
        <p>Adress: xxxxxx</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
      account: "",
      provider: null,
      loginDialogVisible: false,
      accountDrawerVisible: false,
    };
  },

  created() {
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

  .logo {
    width: 200px;
    height: 80px;
    margin-left: 150px;
    cursor: pointer;
    align-items: center;

    p {
      font-size: 40px;
      line-height: 50px;
      margin-top: 15px;
      color: #16952d;
      font-weight: bold;
    }
  }


  .navigation {
    width: 300px;
    height: 80px;
    margin-top: 8px;
    // margin-left: 400px;
    right: -10%;
    display: flex;
    position: relative;

    input {
      width: 300px;
      height: 45px;
      border: 2px solid #5d5d5d;
      border-radius: 10px;
      margin-top: 12px;
    }
  }


  .searchBtn {
    height: 44px;
    width: 60px;
    background-color: #169608;
    border-radius: 12px;

    margin-top: 12px;
    margin-left: 4px;

    .searchIcon {
      height: 44px;
      width: 60px;
      color: #fff;
      margin: auto;
    }

    &:hover {
      background-color: #67C23A;
      color: #FFF;
      cursor: pointer;
    }
  }

  .menu {
    height: 80px;
    display: flex;

    .menuSelection {
      color: #169608;
      width: 150px;
      font-size: 26px;
      line-height: 80px;
      font-weight: bold;
      margin: auto;
      text-align: center;

      &:hover {
        color: #67C23A;
      }
    }
  }


  .Login {
    display: flex;
    margin-left: 200px;
    line-height: 80px;
    font-weight: 700;
    font-size: 26px;
    color: #169608;

    &:hover {
      color: #67C23A;
      cursor: pointer;
    }
  }

}

//底部区域
.bottom {
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
      width: 300px;
      height: 200px;
    }
  }
}
</style>
