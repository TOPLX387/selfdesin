<template>

    <el-card>
        <el-row :gutter="25">
            <el-col :span="7">

                <el-input placeholder="请输入搜索的车牌号" v-model.lazy="queryInfo.searchLpnum" @change="findNewstockinlist">
                </el-input>
            </el-col>

            <el-col :span="4">
                <el-button type="primary" @click="dialogAddVisible = true">货物入库</el-button>
            </el-col>
        </el-row>

        <el-table :data="Newstockinlist" border stripe v-loading="isLoading"
            element-loading-background="rgba(255, 255, 255, .5)"
            element-loading-spinner="el-icon-loading">
            <el-table-column label="序号" type="index" fixed="left"></el-table-column>
            <el-table-column label="车牌号" prop="Lpnum"></el-table-column>
            <el-table-column label="入库件数" prop="Ibpieces"></el-table-column>
            <el-table-column label="入库时间" prop="Whtime"></el-table-column>
            <el-table-column label="入库吨数" prop="Ibton"></el-table-column>
            <el-table-column label="操作" fixed="right">
                <template #default="scope">

                    <el-button type="primary" icon="el-icon-edit" size="mini" circle @click="
                        dialogEditOpen(
                            scope.row.Lpnum,
                            scope.row.Ibpieces,
                            scope.row.Whtime,
                            scope.row.Ibton,
                        )
                    ">
                    </el-button>

                    <el-button type="danger" icon="el-icon-delete" size="mini" circle
                        @click="deleteNewstockin(scope.row.Lpnum)">
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-card>

    <el-dialog title="添加订单" v-model="dialogAddVisible" width="50%">
        <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="100px">
            <el-form-item label="车牌号" prop="Lpnum">
                <el-input v-model="addForm.Lpnum" @blur="test()"></el-input>
            </el-form-item>
            <el-form-item label="入库件数" prop="Ibpieces">
                <el-input v-model="addForm.Ibpieces"></el-input>
            </el-form-item>
            <el-form-item label="入库时间" prop="Whtime">
                <el-date-picker v-model="addForm.Whtime"  value-format="YYYYMMDD"></el-date-picker>
            </el-form-item>
            <el-form-item label="入库吨数" prop="Ibton">
                <el-input v-model="addForm.Ibton"></el-input>
            </el-form-item>
        </el-form>
        <div class="right">
            <el-button type="primary" @click="addNewstockin()">确 定</el-button>
            <el-button @click="dialogAddVisible = false">取 消</el-button>
        </div>
    </el-dialog>


    <el-dialog title="编辑" v-model="dialogEditVisible" width="50%">
        <el-form :model="editForm" :rules="editFormRules" ref="editFormRef" label-width="100px">
            <el-form-item label="车牌号" prop="Lpnum">
                <el-input v-model="editForm.Lpnum"></el-input>
            </el-form-item>
            <el-form-item label="入库件数" prop="Ibpieces">
                <el-input v-model="editForm.Ibpieces"></el-input>
            </el-form-item>
            <el-form-item label="入库时间" prop="Whtime">
                <el-date-picker v-model="editForm.Whtime" value-format="YYYYMMDD" disabled></el-date-picker>
            </el-form-item>
            <el-form-item label="入库吨数" prop="Ibton">
                <el-input v-model="editForm.Ibton"></el-input>
            </el-form-item>
        </el-form>
        <div class="right">
        <el-button type="primary" @click="updateNewstockin">确 定</el-button>
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
                searchLpnum: "",
                page: 1,
                limit: 10,
            },
            Newstockinlist: [],
            total: 0,
            addForm: {
                Lpnum: "",
                Ibpieces: "",
                Whtime: "",
                Ibton: "",
            },
            editForm: {
                Lpnum: "",
                Ibpieces: "",
                Whtime: "",
                Ibton: "",
            },
            addFormRules: {
                Lpnum: [
                    {
                        required: true,
                        message: "请输入车牌号",
                        trigger: "blur",
                    },
                ],
                Ibpieces: [
                    {
                        required: true,
                        message: "请输入入库件数",
                        trigger: "blur",
                    },
                ],
                Whtime: [
                    {
                        required: true,
                        message: "请输入入库时间",
                        trigger: "blur",
                    },
                ],
                Ibton: [
                    {
                        required: true,
                        message: "请输入入库吨数",
                        trigger: "blur",
                    },
                ],
            },
            editFormRules: {
                Lpnum: [
                    {
                        required: false,
                        message: "请输入车牌号",
                        trigger: "blur",
                    },
                ],
                Ibpieces: [
                    {
                        required: false,
                        message: "请输入入库件数",
                        trigger: "blur",
                    },
                ],
                Whtime: [
                    {
                        required: false,
                        message: "请输入入库时间",
                        trigger: "blur",
                    },
                ],
                Ibton: [
                    {
                        required: false,
                        message: "请输入入库吨数",
                        trigger: "blur",
                    },
                ],
            },
        };
    },
    created() {
        this.getNewstockinlist();
    },
    methods: {
        async getNewstockinlist() {
            this.isLoading = true;

            console.log(this.Newstockinlist)
            const { data: res } = await this.$http.get("guobaichuan/route/Newstockin", {
                params: this.queryInfo,
            });
            if (res.code != 20000) {
                this.$message.error("加载出入库信息列表失败");
                this.isLoading = false;
            }
            this.Newstockinlist = res.data.value;
            this.total = res.data.total;
            this.isLoading = false;
        },

        async findNewstockinlist() {
            this.isLoading = true;

            console.log(this.Newstockinlist)
            const { data: res } = await this.$http.get("guobaichuan/route/Newstockins", {
                params: this.queryInfo,
            });
            if (res.code != 20000) {
                this.$message.error("加载信息失败");
                this.isLoading = false;
            }
            this.Newstockinlist = res.data.value;
            this.total = res.data.total;
            this.isLoading = false;
        },



        handleSizeChange(newLimit) {
            this.queryInfo.limit = newLimit;
            this.getNewstockinlist();
        },

        handleCurrentChange(newPage) {
            this.queryInfo.page = newPage;
            this.getNewstockinlist();
        },
        test() {
            console.log(this.addForm.Lpnum);
        },
        async deleteNewstockin(Lpnum) {

            const confirmResult = await this.$confirm(
                "是否确定删除?",
                "提示",
                {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning",
                }
            ).catch((err) => err);

            if (confirmResult != "confirm") {
                return this.$message.info("已取消删除");
            }
            const { data: res } = await this.$http.delete(
                "guobaichuan/route/Newstockin?lpnum=" + Lpnum

            );
            console.log(res);
            console.log(Lpnum)
            if (res.code != 20000) {
                return this.$message.error("删除失败");
            }
            this.$message.success("删除成功");
            this.getNewstockinlist();
            console.log(Lpnum)
        },

        addNewstockin() {
            this.$refs.addFormRef.validate(async (valid) => {
                if (!valid) return;

                console.log(this.addForm);
                const { data: res } = await this.$http.post(
                    "/guobaichuan/route/Newstockin",
                    this.addForm
                );
                if (res.code == 20000) {
                    this.getNewstockinlist();
                    this.dialogAddVisible = false;
                    console.log(this.Newstockinlist)
                    return this.$message.success("添加成功");
                }
                this.$message.error("添加失败");
            });
        },

        updateNewstockin() {
            this.$refs.editFormRef.validate(async (valid) => {
                if (!valid) return;

                console.log(this.Lpnum)
                const { data: res } = await this.$http.put(
                    "/guobaichuan/route/Newstockin",
                    this.editForm
                );
                if (res.code == 20000) {
                    this.getNewstockinlist();
                    this.dialogEditVisible = false;
                    return this.$message.success("编辑成功");
                }
                this.$message.error("编辑失败");
            });
        },
        dialogEditOpen(
            Lpnum,
            Ibpieces,
            Whtime,
            Ibton
        ) {
            this.editForm.Lpnum = String(Lpnum);
            this.editForm.Ibpieces = String(Ibpieces);
            this.editForm.Whtime = String(Whtime);
            this.editForm.Ibton = String(Ibton);
            this.dialogEditVisible = true;
        },
    },
};
</script>

<style>
</style>
