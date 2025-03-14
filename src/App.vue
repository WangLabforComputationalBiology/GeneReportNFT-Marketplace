<template>
  <!-- 头部导航栏 -->
  <header>
    <div class="side" @click="toHome">
      LOGO
    </div>
    <div class="headerWrapper">
      <!-- 路由菜单 -->
      <div class="routers">
        <router-link to="/market" class="routerSelection">Market</router-link>
        <router-link to="/create" class="routerSelection">Create</router-link>
        <router-link to="/drop" class="routerSelection">Drop</router-link>
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
      <!-- 登录路由 -->
      <router-link to="/login" class="routerSelection" style="right:0;">Login</router-link>
    </div>
    <span class="side" @click="toHome">
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
    width: 1400px;
    display: flex;
    margin: 0 auto;

    .navigation {
      margin-left: 300px;
      margin-right: 100px;
      width: 300px;
      display: flex;
      align-items: center;
      margin-top: 8px;

      input {
        width: 300px;
        height: 40px;
        border: 2px solid #5d5d5d;
        border-radius: 10px;
      }

      .searchBtn {
        background-color: #169608;
        border-radius: 12px;
        margin-left: 4px;

        .searchIcon {
          height: 40px;
          width: 60px;
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
      margin: 0 30px;
      font-size: 26px;
      align-content: center;
      font-weight: bold;

      &:hover {
        color: #67C23A;
      }

      &:first-child {
        margin-left: 0;
      }

      // &:last-child {
      //   margin-left: 20px;
      // }
    }


  }
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
      width: 300px;
      height: 200px;
    }
  }
}
</style>
