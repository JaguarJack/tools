<template>
  <div class="container" style="margin-top:10px;">
    <Tabs type="card" @on-click="getName">
      <TabPane label="全部" name="all">
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
      </TabPane>
      <TabPane v-for="tab in tabs" :key="+ tab.Id +''" :label="tab.Name" :name="tab.Code">
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
      </TabPane>
  <!--    <TabPane label="标签一" name="name1" >
        <Card class="card col-sm-3 float-left" v-for="tool in tools">
          <div>
            <div class="text-center">
              <img :src="tool.Icon">
              <h5 slot="title">{{ tool.Title }}</h5>
            </div>
            <div>
             {{ tool.Description }}
            </div>
            <p>
              <Button type="error" size="small" class="float-right">进入</Button>
            </p>
          </div>
        </Card>
      </TabPane>
      <TabPane label="标签二" name="name2">
        <Card class="card col-sm-3 float-left">
          <div>
            <div class="text-center">
              <img src="http://tool.oschina.net/img/logo/apidocs.gif">
              <h5 slot="title">Borderless card</h5>
            </div>
            <div>
              标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一
            </div>
            <p>
              <Button type="error" size="small" class="float-right">进入1</Button>
            </p>
          </div>
        </Card>
      </TabPane>
      <TabPane label="标签三" name="name3">
        <Card class="card col-sm-3 float-left">
          <div>
            <div class="text-center">
              <img src="http://tool.oschina.net/img/logo/apidocs.gif">
              <h5 slot="title">Borderless card</h5>
            </div>
            <div>
              标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一标签一
            </div>
            <p>
              <Button type="error" size="small" class="float-right">进入2</Button>
            </p>
          </div>
        </Card>
      </TabPane>-->
    </Tabs>

  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      tabs: null,
      tools: null
    }
  },
  created () {
    this.http.get(this.host + 'getCategory').then((response) => {
      this.tabs = response.data.data
      console.log(this.tabs)
    })
    this.http.post(this.host + 'getTools', {name: 'all'}).then((response) => {
      this.tools = response.data.data
      console.log(this.tools)
    })
  },
  methods: {
    getName (name) {
      this.http.post(this.host + 'getTools', {name: name}).then((response) => {
        this.tools = response.data.data
        console.log(this.tools)
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
