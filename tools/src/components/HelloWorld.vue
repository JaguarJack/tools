<template>
  <div class="container" style="margin-top:10px;">
    <Tabs type="card" @click="getName()" name="name1">
      <TabPane label="标签一">
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
              <Button type="error" size="small" class="float-right">进入</Button>
            </p>
          </div>
        </Card>
      </TabPane>
      <TabPane label="标签二">
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
      <TabPane label="标签三">
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
      </TabPane>
    </Tabs>

  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js Apps'
    }
  },
  methods: {
    getName (name) {
      console.log(name)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
