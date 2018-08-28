<template>
  <div class="container" style="margin-top:10px;">
    <Card v-for="tool in tools" :key="+tool.Id+''" class="card col-sm-3 float-left" >
      <div>
        <div class="text-center">
          <img :src="tool.Icon">
          <h5 slot="title">{{ tool.Title }}</h5>
        </div>
        <div>
          {{ tool.Description }}
        </div>
        <p>
          <router-link :to="'/'+ tool.Code"><Button type="error" size="small" class="float-right">进入</Button></router-link>
        </p>
      </div>
    </Card>
  </div>
</template>

<script>
export default {
  name: 'index',
  data () {
    return {
      tools: null
    }
  },
  created () {
    this.http.post(this.host + 'getTools', {name: 'all'}).then((response) => {
      this.tools = response.data.data
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
