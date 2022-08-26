<template>
  <el-backtop :bottom="60"></el-backtop>
  <el-container class="home-container">
    <el-header>
      <div>
        <div style="padding: 0 30px">
        港口堆存费管理系统
      </div>
      <el-menu background-color="#282828" text-color="#fff" active-text-color="#409eff" unique-opened mode="horizontal"
          :collapse="isCollapse" :collapse-transition="false" :router="true" :default-active="$route.path">
          <el-menu-item :index="it.MenuAddress" v-for="it in menuList" :key="it.MenuId">
            <i class="el-setting"></i>
            <template #title>{{ it.MenuName }}</template>
          </el-menu-item>
        </el-menu>
      </div>
      <el-button type="info" @click="logout">退出</el-button>
    </el-header>
    <el-container>
        <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>
<script>
export default {
  data() {
    return {

      menuList: [],
      isCollapse: false,
      dialogVisible: false,
      username: "",
      level: 0,
      remarks: "",
    };
  },
  created() {
    this.username = window.sessionStorage.getItem("username");
    this.level = window.sessionStorage.getItem("level");
    this.remarks = window.sessionStorage.getItem("remarks");
    this.getMenuList();
  },
  methods: {
    logout() {
      window.sessionStorage.clear();
      this.$router.push("/login");
    },

    async getMenuList() {
      const { data: res } = await this.$http.get(
        "guobaichuan/route/menu?level=" + this.level
      );

      if (res.code != 20000) return this.$message.error("菜单加载失败");
      this.menuList = res.data.menus;

    },
  },
};
</script>

<style lang="less" scoped>
.el-header {
  background-color: #000000;
  display: flex;
  justify-content: space-between;
  padding-left: 0%;
  align-items: center;
  color: #fff;
  font-size: 20px;

  >div {

    display: flex;
    align-items: center;

    span {
      margin-left: 15px;
    }
  }
}

.el-aside {
  background-color: #333744;

  .el-menu {
    border-right: none;
  }
}
.el-main {
  background-color: #eaedf1;
  //background-image: url(../assets/login.jpg);
}

.home-container {
  height: 100%;

}

.logo_img {
  width: 20%;
  height: 100%;
}

.iconfont {
  margin-right: 10px;
}

.toggle-button {
  background-color: #4a5064;
  font-size: 10px;
  line-height: 24px;
  color: #fff;
  text-align: center;
  letter-spacing: 0.2em;
  cursor: pointer;
}
</style>
