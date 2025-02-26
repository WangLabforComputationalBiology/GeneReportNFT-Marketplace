<template>
  <!-- 头部导航栏 -->
  <div class="header">
    <span class="logo" @click="toHome">
      <p>LOGO</p>
    </span>

    <div class="menu">
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

    <router-link to="/loginPage"><span class="accountAddress"><i class="el-icon-user"></i>Account: {{ accountdisplay }}
      </span></router-link>
  </div>


  <router-view />

  <!-- 尾部导航栏 -->
  <div class="bottom">
    <div class="inBottom">
      <div class="communication">
        <p>Need some help? Please contact us.</p> <br>
       
        <p>WangLab @SZTU</p>
        <p>Email: 11111@gmail.com</p>
        <p>Adress: xxxxxx</p>
      </div>
    </div>

  </div>
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
a {
  text-decoration: none;
}

.header {
  width: 1400px;
  height: 80px;

  display: flex;
  position: sticky;
  top: 0;
  transition: background-color 0.5s;

  .logo {
    width: 200px;
    height: 80px;
    margin: 0 80px;
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
    margin-left: 400px;
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
    font-size: 25px;
    line-height: 80px;
    font-weight: bold;
    margin: auto;
    text-align: center;

    &:hover {
      color: #67C23A;
    }
  }
}


.accountAddress {
  display: flex;
  line-height: 80px;
  font-size: 18px;
  color: #169608;
  cursor: pointer;

  .el-icon-user {
    font-size: 24px;
    margin: 0 5px 0 100px;
    line-height: 80px;
  }
}

// 登录弹窗
:deep(.custom-dialog) {
  border-radius: 10px !important;
  width: 20%;
  margin-top: 5%;
  box-shadow: 0 0 6px #DCDFE6;

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
        box-shadow: 0 0 6px #DCDFE6;
        transition: 200ms;
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


.bottom {
  width: 100%;
  height: 300px;
  background-color: #169608;
  margin: 0 auto;
  display: flex;

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
