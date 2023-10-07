<template>
  <div class="apidata">
    <span>{{ token }}</span>
  </div>
</template>

<script>
export default {
  name: "Getdata",
  data() {
    return {
      token: "",
    };
  },
  methods: {
    async initToken() {
      // 请求回来的数据在页面要用到，必须转存到data中
      // const { data: res1 } = await this.$http.get("/api/dept");
      // const { token: res } = await request.get("/api/dept");
      const { token: res } = await this.$http.get(
        "http://localhost:3000/info",
      );
      console.log(res);
    },
  },
  created() {
    this.initToken();
  },
};
</script>

<style></style>
