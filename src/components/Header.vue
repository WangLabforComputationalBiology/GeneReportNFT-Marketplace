<template>
    
 

</template>
  
<script>
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
</script>

<style>
@import '../assets/Header.css';

</style>