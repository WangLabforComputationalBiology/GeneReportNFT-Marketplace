<template>

 
    <!-- 头部导航栏 -->
    <div class="header">

  <span id="logo">
    <p>LOGO</p>
  </span>
  <ul>
    <!-- 铸造 -->
    <li id="create"><a href="#"><p>Create</p></a></li>

    <!-- 投放 -->
    <li id="drop"><a href="#"><p>Drop</p></a></li>

    <!-- 统计 -->
    <li id="stats"><a href="#"><p>Stats</p></a></li>
  </ul>

  <!-- 搜索框 -->
  <span id="navigation">
    <input id="navigationInput" type="text" placeholder="press '/' to search...">
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
    <router-link to="/test1" @click="test"> test1 </router-link>
    <router-link to="/test2" > test2 </router-link>

  </div>

  <div id="display">
    <router-view></router-view>
  </div>
  

  <!-- 尾部导航栏
  <div class="bottom">
    <p>bottom area</p>
  </div> -->

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

#display{
  width: 200px;
  height: 200px;
  border: 1px solid #fff;
}

</style>
