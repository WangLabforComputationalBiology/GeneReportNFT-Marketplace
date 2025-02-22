<template>
  <!-- 头部导航栏 -->
  <div class="header">
    <span id="logo" @click="toHome">
      <p>LOGO</p>
    </span>

    <div class="menu">
      <router-link to="/create" class="menuSelection">Create</router-link>
      <router-link to="/drop" class="menuSelection">Drop</router-link>
      <router-link to="/stats" class="menuSelection">Stats</router-link>
    </div>

    <!-- 搜索框 -->
    <span id="navigation">
      <input id="navigationInput" type="text" placeholder="search..." />
      <div type="primary" class="searchBtn">
        <el-icon class="searchIcon" size="large">
          <Search />
        </el-icon>
      </div>
    </span>

    <!-- 登录界面 -->
    <el-dialog v-model="loginDialogVisible" title="Login" center class="custom-dialog">
      <div class="dialogBody">
        <p>Try this:</p>
        <div class="useMeta" @click="connectWallet">
          <img class="icon" src="./icons/metalogo.png" />
          <span class="meta">
            <p>MetaMask</p>
          </span>
        </div>
        <div class="loginNotice">Notice: Please use Chrome or Firefox and make sure you have downloaded MetaMask
          as an extension. </div>
      </div>
    </el-dialog>

    <span class="loginBTN" @click="isVisible">
      Login
    </span>

    <span class="accountAddress">Account: {{ accountdisplay }} </span>
  </div>


  <router-view />

  <!-- 尾部导航栏
  <div class="bottom">
    <p>bottom area</p>
  </div> -->
</template>

<script>
import { ref } from 'vue';
import { ethers } from "ethers";

export default {
  name: "App",
  data() {
    return {
      account: "",
      provider: null,
      loginDialogVisible: false,
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
      this.$router.push("/");
    },

    //控制登录弹窗
    isVisible() {
      this.loginDialogVisible = true;
    },

    // MetaMask连接并获取账户
    async connectWallet() {
      if (typeof window.ethereum !== 'undefined') {
        try {
          const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
          this.account = accounts[0];
          console.log('Connected account:', this.account);
        } catch (error) {
          console.error('User denied account access or error occurred:', error);
        }
      } else {
        console.log('MetaMask is not installed');
      }

      if (this.account) {
        this.$message({
          message: 'Wallet Connected successfully!',
          type: 'success',
          duration: 2000,
        });
        this.loginDialogVisible = false;
      }
    },





  },
}

</script>


<style lang="scss" scoped>
@import "./assets/main.css";

a {
  text-decoration: none;
}

.header {
  background: #fff;
  width: 100%;
  height: 80px;

  display: flex;
  position: sticky;
  top: 0;
  transition: background-color 0.5s;
  border-bottom: 1px solid #EBEEF5;
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
  width: 500px;
  height: 80px;
  display: flex;

  .menuSelection {
    background-color: #ffffff;
    color: #169608;
    width: 140px;
    font-size: 25px;
    line-height: 80px;
    font-weight: bold;
    margin: auto;
    text-align: center;
    border-bottom: 1px solid #EBEEF5;

    &:hover {
      background-color: #fffffff1;
      color: #67C23A;
    }
  }
}

.loginBTN {
  height: 44px;
  width: 60px;
  background-color: #169608;
  border-radius: 12px;
  cursor: pointer;
  color: #fff;
  font-size: 16px;
  text-align: center;
  line-height: 44px;

  margin: 20px 0 0 200px;
  float: right;

  &:hover {
    background-color: #67C23A;
    color: #FFF;

  }

}

.accountAddress {
  margin: 28px 0 0 40px;
  font-size: 18px;
  color: #169608;
  cursor: pointer;
}

// 登录弹窗
::v-deep .custom-dialog {
  border-radius: 10px !important;
  width: 20%;
  margin-top: 5%;

  .dialogBody {
    height: 300px;

    .useMeta {
      width: 100%;
      height: 60px;
      margin-top: 30px;
      font-size: 20px;
      font-weight: bold;
      color: #67C23A;
      border: grey 1px solid;
      border-radius: 10px;
      display: flex;

      .icon {
        margin: 10px;
        width: 40px;
        height: 40px
      }

      .meta {

        p {
          line-height: 60px;
          margin-left: 20px;
          color: rgb(29, 29, 29);
        }
      }



      &:hover {
        cursor: pointer;
      }

    }


    .loginNotice {
      width: 100%;
      height: 50px;
      margin-top: 100px;
    }

    p {
      font-size: 16px;
      font-weight: bold;
      color: #169608;
    }



  }
}








.mainBody {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;

}
</style>
