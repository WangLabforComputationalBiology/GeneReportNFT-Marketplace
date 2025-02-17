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

    <el-dialog v-model="loginDialogVisible" title="Login" width="20%" center class="custom-dialog">
      <div class="dialogBody">
        <p>Try this:</p>
        <div class="useMeta">
          <img class="icon" src="./icons/metalogo.png" />
          <span class="meta">
            <p>MetaMask</p>
          </span>
        </div>
      </div>

    </el-dialog>

    <span class="loginBTN" @click="isVisible">
      Login
    </span>
  </div>


  <router-view />

  <!-- 尾部导航栏 -->
  <div class="bottom">
    <p>bottom area</p>
  </div>
</template>

<script>
import { ref } from 'vue';

export default {
  name: "App",
  data() {
    return {
      loginDialogVisible: false,
    };
  },
  methods: {
    toHome() {
      this.$router.push("/");
    },
    isVisible() {
      this.loginDialogVisible = true;
    },

  },


};
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
  border-bottom: 1px solid #F2F6FC;
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
    border-bottom: 1px solid #F2F6FC;

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


.custom-dialog {
  .el-dialog__body {
    margin-top: 10%;
  }

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
