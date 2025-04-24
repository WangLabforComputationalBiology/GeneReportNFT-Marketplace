<template>
    <div class="footer">
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

<style lang="scss" scoped>
.footer {
    width: 99vw;
    // max-width: 99vw;
    height: 300px;
    background-color: #169608;
    margin: auto;
    display: flex;

    //底部中心内容
    .inBottom {
        min-width: 1400px;

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