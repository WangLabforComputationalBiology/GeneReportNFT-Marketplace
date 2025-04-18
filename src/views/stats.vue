<template>
    <body>
        <div class="banner">
            <h1 class="banner-title">Stats</h1>
        </div>

        <div class="statsPage">
        </div>  
    </body>
</template>

<script>
export default {
    name: "stats"
}
</script>

<style lang="scss" scoped>

body{
margin: auto;
width: 80vw;
min-width: 1400px;
background-color: #fff;
height: calc(100vh - 80px);
}

.banner-title{
font-size: 70px;
color: #67C23A;
}

.statsPage{
height: 700px;
}
</style>