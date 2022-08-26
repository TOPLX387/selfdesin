<template>
  <div class="login_container">
    <div class="login_box">
      <div class="avatar_box">
      </div>
      <center>
        <h1>港口堆存费管理系统</h1>
      </center>
      <el-form ref="loginFormRef" :model="loginForm" :rules="loginRules" class="login_form" label-width="0px">
        <el-form-item prop="用户名">
          <el-input v-model="loginForm.username" prefix-icon="el-icon-user"></el-input>
        </el-form-item>
        <el-form-item prop="密码">
          <el-input v-model="loginForm.password" prefix-icon="el-icon-lock" type="password"></el-input>
        </el-form-item>
        <el-form-item class="btns">
          <el-button @click="login"   type="primary">登录</el-button>
          <el-button @click="resetLoginForm" type="info">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loginForm: {
        username: "",
        password: "",
      },
      loginRules: {
        username: [
          {
            required: true,
            message: "请输入用户名",
            trigger: "blur",
          },
          {
            min: 5,
            max: 12,
            message: "长度在 5 到 12 个字符",
            trigger: "blur",
          },
        ],
        password: [
          {
            required: true,
            message: "请输入密码",
            trigger: "blur",
          },
          {
            min: 5,
            max: 12,
            message: "密码为 5~12 位",
            trigger: "blur",
          },
        ],
      },
    };
  },
  methods: {
    resetLoginForm() {
      this.$refs.loginFormRef.resetFields();
    },
    login() {
      this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return;
        const { data: res } = await this.$http.post(
          "guobaichuan/route/login",
          this.loginForm
        );
        if (res.code == 30000) {
          this.$message.error("没有此用户");
          return;
        }
        if (res.code == 30001) {
          this.$message.error("密码错误");
          return;
        }
        if (res.code == 20000) {
          this.$message.success("登录成功");
          window.sessionStorage.setItem("username", res.data.user.LogintoName);
          window.sessionStorage.setItem("level", res.data.user.LogintoLevel);
          window.sessionStorage.setItem("remarks", res.data.user.LogintoRemarks);
          this.$router.push({ path: "/home" });
          return;
        }
        this.$message.error("未知错误");
      });
    },
  },
};
</script>

<style lang="less" scoped>
.login_container {
  background-color: #f6f8f6;
  height: 100%;
  background-size: 100%;
  background-image: url(../assets/5.jpeg);

}

.login_box {
  width: 350px;
  height: 300px;
  border-radius: 3px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);



  .btns {
    display: flex;
    justify-content: center;
  }

  .login_form {
    position: absolute;
    bottom: 0%;
    width: 100%;
    padding: 0 10px;
    box-sizing: border-box;
  }
}
</style>
