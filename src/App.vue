<template>

 
    <!-- 头部导航栏 -->
    <div class="header">

  <span id="logo">
    <p>LOGO</p>
  </span>
  <div>
    <el-menu mode="horizontal" class="menu" active-text-color="#169608">
      <el-menu-item index="1" class="menuSelection"><router-link to="/create">Create</router-link></el-menu-item>
      <el-menu-item index="2" class="menuSelection"><router-link to="/drop" >Drop</router-link></el-menu-item>
      <el-menu-item index="3" class="menuSelection"><router-link to="/stats" >Stats</router-link></el-menu-item>
    </el-menu>
  </div>
  

  <!-- 搜索框 -->
  <span id="navigation">
    <input id="navigationInput" type="text" placeholder="press '/' to search...">
    <el-button type="primary" class="searchBtn">
            <el-icon ><Search /></el-icon>
        </el-button>
    <!-- <span id="naviButtom"></span> -->
  </span>
  </div>

  <!-- 侧边栏 -->
  <span class="slider">
    <div id="profilePic"> <p>LOGIN</p> </div>
    <div id="userName"> <p>#userName</p> </div>

    <!-- 返回顶部按钮 -->
    <button id="backToTopBTN" @click="backToTop"><p>Back to top ↑</p></button>

   </span>


  

  <!-- 内容主体 -->
  <div class="main">
    <router-view/>
  </div>

  <!-- 尾部导航栏 -->
  <div class="bottom">
    <p>bottom area</p>
  </div>

</template>



<script setup>
  
function test(){
  console.log("1");
}

// header导航栏透明切换方法
document.addEventListener('DOMContentLoaded', function() {
  var headerRolling = document.querySelector('.header');
  window.addEventListener('scroll', function() {
    var scrollHeight = window.pageYOffset || document.documentElement.scrollTop;
    if (scrollHeight > 50) { // 页面滚动超过50px时改变导航栏样式
      headerRolling.style.backgroundColor = 'rgba(255, 255, 255, 0.5)'; // 设置为透明
      headerRolling.style.backdropFilter = 'blur(10px)';    
    } else {
      headerRolling.style.backgroundColor = 'rgba(255, 255, 255, 1) '; // 恢复不透明
      headerRolling.style.backdropFilter = 'none';
    }
  });

 
});



//返回顶部方法
function backToTop() {
    document.body.scrollTop = 0;
    document.documentElement.scrollTop = 0;
}

document.addEventListener('keydown', function(event) {
  /*  检查按下的键是否是“/”且搜索框当前不可见 */
  if (event.key === '/') {
    event.preventDefault(); 
     /* 阻止默认行为，如输入斜杠字符 */

    /* // 显示搜索框 */
    document.getElementById('navigationInput').focus();
     /* // 将焦点设置到搜索框 */
  }
});
</script>




<style scoped>
@import './assets/main.css';

a{
  text-decoration: none; 
}

.searchBtn{
  height: 44px;
  width: 60px;
  background-color:#169608;
  border-radius: 12px;

  margin-top: 12px;
  margin-left: 4px;
}
.searchBtn:hover{
  background-color: #fff;
  color: #169608;
}

.menu{
  width: 500px;
  height: 80px;
}

.menu:hover {
  /* background-color: #2e1e1e; */
  color: #169608;
}
.menuSelection {
  background-color: #ffffff;
  color: #169608;
  width: 140px;
  font-size: 25px;
  font-weight: bold;

}

.menuSelection:hover{
  background-color: #d17070;
}

</style>
