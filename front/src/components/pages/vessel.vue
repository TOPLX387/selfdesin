<template>
  <!-- 卡片视图区 -->
  <el-card>
    <el-row :gutter="25">
      <el-col :span="7">
        <!-- 搜索添加 -->
        <el-input placeholder="请输入货代公司名称" v-model.lazy="queryInfo.searchLadingnum" @change="getLadingList">
        </el-input>
      </el-col>

      <el-col :span="4">
        <el-button type="primary" round @click="dialogAddVisible = true">添加信息</el-button>
      </el-col>
    </el-row>
    <!-- 列表 -->
    <el-table :data="vesselList" border stripe v-loading="isLoading"
      element-loading-background="rgba(255, 255, 255, .5)" element-loading-text="加载中，请稍后..."
      element-loading-spinner="el-icon-loading">
      <el-table-column label="序号" type="index" fixed="left"></el-table-column>
      <el-table-column label="货代公司名称" prop="Ladingnum"></el-table-column>
      <el-table-column label="货代公司编号" prop="Cname"></el-table-column>
      <el-table-column label="提单号" prop="Destination"></el-table-column>
      <el-table-column label="操作" fixed="right">
        <template #default="scope">
          <!-- 修改 -->
          <el-button type="primary" icon="el-icon-edit" size="mini" circle @click="
            dialogEditOpen(
              scope.row.Ladingnum,
              scope.row.Cname,
              scope.row.Destination
            )
          ">
          </el-button>
          <!-- 删除 -->
          <el-button type="danger" icon="el-icon-delete" size="mini" circle @click="deleteLading(scope.row.Cname)">
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="queryInfo.page"
      :page-sizes="[1, 2, 10, 100]" :page-size="queryInfo.limit" layout="total, sizes, prev, pager, next, jumper"
      :total="total"></el-pagination>
  </el-card>

  <el-dialog title="添加信息" v-model="dialogAddVisible" width="50%">
    <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="150px">
      <el-form-item label="货代公司名称" prop="Ladingnum">
        <el-input v-model="addForm.Ladingnum"></el-input>
      </el-form-item>
      <el-form-item label="货代公司编号" prop="Cname">
        <el-input v-model="addForm.Cname"></el-input>
      </el-form-item>
       <el-form-item label="提单号" prop="Destination">
        <el-input v-model="addForm.Destination"></el-input>
      </el-form-item>
    </el-form>
     <div class="right">
    <el-button type="primary" @click="addLading()">确 定</el-button>
    <el-button @click="dialogAddVisible = false">取 消</el-button>
    </div>
  </el-dialog>

  <!-- 编辑对话框 -->
  <el-dialog title="编辑信息" v-model="dialogEditVisible" width="50%">
    <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="150px">
      <el-form-item label="货代公司名称" prop="Ladingnum">
        <el-input v-model="editForm.Ladingnum"></el-input>
      </el-form-item>
      <el-form-item label="货代公司编号" prop="Cname">
        <el-input v-model="editForm.Cname"></el-input>
      </el-form-item>
      <el-form-item label="提单号" prop="Destination">
        <el-input v-model="editForm.Destination"></el-input>
      </el-form-item>
    </el-form>
     <div class="right">
    <el-button type="primary" @click="updateLading">确 定</el-button>
    <el-button @click="dialogEditVisible = false">取 消</el-button>
    </div>
  </el-dialog>
</template>

<script>
export default {
  data() {
    return {
      isLoading: false,
      dialogAddVisible: false,
      dialogEditVisible: false,
      queryInfo: {
        searchLadingnum: "",
        page: 1,
        limit: 10,
      },
      vesselList: [], // 信息列表
      total: 0, // 最大数据记录
      addForm: {
        Ladingnum: "",
        Cname: "",
        Destination:"",
      },
      editForm: {
        Ladingnum: "",
        Cname: "",
        Destination:"",
      },
      addFormRules: {
        Ladingnum: [
          {
            required: true,
            message: "请输入货代公司名称称",
            trigger: "blur",
          },
        ],
        Cname: [
          {
            required: true,
            message: "请输入货代公司编号",
            trigger: "blur",
          },
        ],
        Destination: [
          {
            required: true,
            message: "请输入提单号",
            trigger: "blur",
          },
        ],
      },
      editFormRules: {
        Ladingnum: [
          {
            required: true,
            message: "请输入货代公司名称称",
            trigger: "blur",
          },
        ],
        Cname: [
          {
            required: true,
            message: "请输入货代公司编号",
            trigger: "blur",
          },
        ],
        Destination: [
          {
            required: true,
            message: "请输入提单号",
            trigger: "blur",
          },
        ],
      },
    };
  },
  created() {
    this.getLadingList();
  },
  methods: {
    async getLadingList() {
      this.isLoading = true;
      // 调用post请求
      // console.log(this.vesselList)
      const { data: res } = await this.$http.get("guobaichuan/route/vessel", {
        params: this.queryInfo,
      });
      if (res.code != 20000) {
        this.$message.error("加载信息列表失败");
        this.isLoading = false;
      }
      this.vesselList = res.data.value; // 将返回数据赋值
      this.total = res.data.total; // 总个数
      this.isLoading = false;
    },
    // 监听pageSize改变的事件
    handleSizeChange(newLimit) {
      this.queryInfo.limit = newLimit;
      this.getLadingList(); // 数据发生改变重新申请数据
    },
    // 监听pageNum改变的事件
    handleCurrentChange(newPage) {
      this.queryInfo.page = newPage;
      this.getLadingList(); // 数据发生改变重新申请数据
    },
    async deleteLading(Cname) {
      // 弹框
      console.log("货代公司编号")
      console.log(Cname)
      const confirmResult = await this.$confirm(
        console.log(this.vesselList),
        "是否确定删除",
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
        }
      ).catch((err) => err);
      // 成功删除为confirm 取消为 cancel
      if (confirmResult != "confirm") {
        return this.$message.info("已取消删除");
      }
      const { data: res } = await this.$http.delete(
        "guobaichuan/route/vessel?cname=" + Cname
      );
      console.log("删除")
      console.log("guobaichuan/route/vessel?cno=" + Cname)
      if (res.code != 20000) {
        console.log(this.Cname)
        return this.$message.error("删除失败");
      }
      this.$message.success("删除成功");
      this.getLadingList();
    },
    // 添加信息
    addLading() {
      this.$refs.addFormRef.validate(async (valid) => {
        if (!valid) return;
        // 发起请求
        const { data: res } = await this.$http.post(
          "/guobaichuan/route/vessel",
          this.addForm
        );
        if (res.code == 20000) {
          this.getLadingList();
          this.dialogAddVisible = false;
          return this.$message.success("添加成功");
        }
        this.$message.error("添加失败");
      });
    },
    // 编辑
    updateLading() {
      this.$refs.editFormRef.validate(async (valid) => {
        if (!valid) return;
        // 发起请求
        const { data: res } = await this.$http.put(
          "/guobaichuan/route/vessel",
          this.editForm
        );
        if (res.code == 20000) {
          this.getLadingList();
          this.dialogEditVisible = false;
          return this.$message.success("编辑成功");
        }
        this.$message.error("编辑失败");
      });
    },
    dialogEditOpen(
      Ladingnum,
      Cname,
      Destination,
    ) {
      this.editForm.Ladingnum = String(Ladingnum);
      this.editForm.Cname = String(Cname);
      this.editForm.Destination = String(Destination);
      this.dialogEditVisible = true;
    },
  },
};
</script>

<style>
</style>
