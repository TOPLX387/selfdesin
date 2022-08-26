<template>

    <el-card>
        <el-row :gutter="25">
            <el-col :span="7">

                <el-input placeholder="请输入搜索的订单号" v-model.lazy="queryInfo.searchOrdernum" @change="findNewstockoutlist">
                </el-input>
            </el-col>

            <el-col :span="4">
<!--                <el-button type="primary" round @click="dialogAddVisible = true">出库</el-button>-->
              <el-button type="primary"  @click="dialogAddVisible = true">货物出库</el-button>
            </el-col>
        </el-row>

        <el-table :data="Newstockoutlist" border stripe v-loading="isLoading"
            element-loading-background="rgba(255, 255, 255, .5)"
            element-loading-spinner="el-icon-loading">
            <el-table-column label="序号" type="index" fixed="left"></el-table-column>
            <el-table-column label="订单号" prop="Ordernum"></el-table-column>
            <el-table-column label="货代公司名称" prop="Ladingnum"></el-table-column>
            <el-table-column label="出库件数" prop="Outpieces"></el-table-column>
            <el-table-column label="出库时间" prop="Detime"></el-table-column>
            <el-table-column label="出库吨数" prop="Obton"></el-table-column>
            <el-table-column label="堆存费" prop="Storagefee"></el-table-column>

            <el-table-column label="操作" fixed="right">
                <template #default="scope">

                    <el-button type="primary" icon="el-icon-edit" size="mini" circle @click="
                        dialogEditOpen(
                            scope.row.Ordernum,
                            scope.row.Ladingnum,
                            scope.row.Outpieces,
                            scope.row.Detime,
                            scope.row.Obton,
                            scope.row.Storagefee,

                        )
                    ">
                    </el-button>

                    <el-button type="danger" icon="el-icon-delete" size="mini" circle
                        @click="deleteNewstockout(scope.row.Ordernum)">
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>

    <el-dialog title="添加订单" v-model="dialogAddVisible" width="50%">
        <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="100px">
            <el-form-item label="订单号" prop="Ordernum">
                <el-input v-model="addForm.Ordernum" @blur="test()"></el-input>
            </el-form-item>
            <el-form-item label="公司名称" prop="Ladingnum">
                <el-input v-model="addForm.Ladingnum"></el-input>
            </el-form-item>
            <el-form-item label="出库件数" prop="Outpieces">
                <el-input v-model="addForm.Outpieces"></el-input>
            </el-form-item>
            <el-form-item label="出库时间" prop="Detime">
                <el-date-picker v-model="addForm.Detime"  value-format="YYYYMMDD"></el-date-picker>
            </el-form-item>
            <el-form-item label="出库吨数" prop="Obton">
                <el-input v-model="addForm.Obton"></el-input>
            </el-form-item>
        </el-form>
        <div class="right">
        <el-button type="primary" @click="addNewstockout()">确 定</el-button>
        <el-button @click="dialogAddVisible = false">取 消</el-button>
        </div>
    </el-dialog>


    <el-dialog title="编辑" v-model="dialogEditVisible" width="50%">
        <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="100px">
            <el-form-item label="订单号" prop="Ordernum">
                <el-input v-model="editForm.Ordernum" disabled></el-input>
            </el-form-item>
            <el-form-item label="公司名称" prop="Ladingnum">
                <el-input v-model="editForm.Ladingnum"></el-input>
            </el-form-item>
            <el-form-item label="出库件数" prop="Outpieces">
                <el-input v-model="editForm.Outpieces"></el-input>
            </el-form-item>
            <el-form-item label="出库时间" prop="Detime">
                <el-date-picker v-model="editForm.Detime"  value-format="YYYYMMDD" disabled></el-date-picker>
            </el-form-item>
            <el-form-item label="出库吨数" prop="Obton">
                <el-input v-model="editForm.Obton"></el-input>
            </el-form-item>
        </el-form>
        <div class="right">
        <el-button type="primary" @click="updateNewstockout">确 定</el-button>
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
                searchOrdernum: "",
                page: 1,
                limit: 10,
            },
            Newstockoutlist: [],
            total: 0,
            addForm: {
                Ordernum: "",
                Ladingnum: "",
                Outpieces: "",
                Detime: "",
                // Storagefee: "",
                Obton: "",
            },
            editForm: {
                Ordernum: "",
                Ladingnum: "",
                Outpieces: "",
                Detime: "",
                // Storagefee: "",
                Obton: "",
            },
            addFormRules: {
                Ordernum: [
                    {
                        required: true,
                        message: "请输入订单号",
                        trigger: "blur",
                    },
                ],
                Ladingnum: [
                    {
                        required: true,
                        message: "请输入公司名称",
                        trigger: "blur",
                    },
                ],
                Outpieces: [
                    {
                        required: true,
                        message: "请输入出库件数",
                        trigger: "blur",
                    },
                ],
                Detime: [
                    {
                        required: true,
                        message: "请输入出库时间",
                        trigger: "blur",
                    },
                ],
                Obton: [
                    {
                        required: true,
                        message: "请输入出库吨数",
                        trigger: "blur",
                    },
                ],
            },
            editFormRules: {
                Ordernum: [
                    {
                        required: false,
                        message: "请输入订单号",
                        trigger: "blur",
                    },
                ],
                Ladingnum: [
                    {
                        required: false,
                        message: "请输入公司名称",
                        trigger: "blur",
                    },
                ],
                Outpieces: [
                    {
                        required: false,
                        message: "请输入出库件数",
                        trigger: "blur",
                    },
                ],
                Detime: [
                    {
                        required: false,
                        message: "请输入出库时间",
                        trigger: "blur",
                    },
                ],
                Obton: [
                    {
                        required: false,
                        message: "请输入出库吨数",
                        trigger: "blur",
                    },
                ],
            },
        };
    },
    created() {
        this.getNewstockoutlist();
    },
    methods: {
        async getNewstockoutlist() {
            this.isLoading = true;

            console.log(this.Newstockoutlist)
            const { data: res } = await this.$http.get("guobaichuan/route/Newstockout", {
                params: this.queryInfo,
            });
            if (res.code != 20000) {
                this.$message.error("加载出入库信息列表失败");
                this.isLoading = false;
            }
            this.Newstockoutlist = res.data.value;
            this.total = res.data.total;
            this.isLoading = false;
        },

         async findNewstockoutlist() {
            this.isLoading = true;

            console.log(this.Newstockoutlist)
            const { data: res } = await this.$http.get("guobaichuan/route/Newstockouts", {
                params: this.queryInfo,
            });
            if (res.code != 20000) {
                this.$message.error("加载信息失败");
                this.isLoading = false;
            }
            this.Newstockoutlist = res.data.value;
            this.total = res.data.total;
            this.isLoading = false;
        },

        handleSizeChange(newLimit) {
            this.queryInfo.limit = newLimit;
            this.getNewstockoutlist();
        },

        handleCurrentChange(newPage) {
            this.queryInfo.page = newPage;
            this.getNewstockoutlist();
        },
        test() {
            console.log(this.addForm.Ordernum);
        },
        async deleteNewstockout(Ordernum) {

            const confirmResult = await this.$confirm(
                "是否确定删除",
                "提示",
                {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning",
                }
            ).catch((err) => err);

            if (confirmResult != "confirm") {
                return this.$message.info("已撤消删除");
            }
            const { data: res } = await this.$http.delete(
                "guobaichuan/route/Newstockout?ordernum=" + Ordernum

            );
            console.log(res);
            console.log(Ordernum)
            if (res.code != 20000) {
                return this.$message.error("删除失败");
            }
            this.$message.success("删除成功");
            this.getNewstockoutlist();
            console.log(Ordernum)
        },

        addNewstockout() {
            this.$refs.addFormRef.validate(async (valid) => {
                if (!valid) return;

                console.log(this.addForm);
                const { data: res } = await this.$http.post(
                    "/guobaichuan/route/Newstockout",
                    this.addForm
                );
                if (res.code == 20000) {
                    this.getNewstockoutlist();
                    this.dialogAddVisible = false;
                    console.log(this.Newstockoutlist)
                    return this.$message.success("添加成功");
                }
                this.$message.error("添加失败");
            });
        },

        updateNewstockout() {
            this.$refs.editFormRef.validate(async (valid) => {
                if (!valid) return;

                console.log(this.Ordernum)
                const { data: res } = await this.$http.put(
                    "/guobaichuan/route/Newstockout",
                    this.editForm
                );
                if (res.code == 20000) {
                    this.getNewstockoutlist();
                    this.dialogEditVisible = false;
                    return this.$message.success("编辑成功");
                }
                this.$message.error("编辑失败");
            });
        },
        dialogEditOpen(
            Ordernum,
            Ladingnum,
            Outpieces,
            Detime,
            Obton
        ) {
            this.editForm.Ordernum = String(Ordernum);
            this.editForm.Ladingnum = String(Ladingnum);
            this.editForm.Outpieces = String(Outpieces);
            this.editForm.Detime = String(Detime);
            this.editForm.Obton = String(Obton);
            this.dialogEditVisible = true;
        },
    },
};
</script>

<style>
</style>
