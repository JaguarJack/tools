<template>
  <div id="app">
    <nav class="navbar navbar-expand-sm bg-dark navbar-dark">
      <div class="container">
        <a class="navbar-brand" href="#">在线工具</a>
      </div>
    </nav>
    <div class="container">
      <div class="layout">
        <Layout>
          <Header>
            <Menu mode="horizontal" active-name="all">
              <MenuItem  v-for="tab in tabs" :key="+ tab.Id +''" :name="tab.Code">
                <router-link :to="{ name: 'index', params: { code: tab.Code  }}">{{ tab.Name }}</router-link>
              </MenuItem>
            </Menu>
          </Header>
          <Layout>
            <Content :style="{padding: '24px 0', minHeight: '580px', background: '#fff'}">
              <Layout>
                <Content :style="{padding: '24px', minHeight: '580px', background: '#fff'}">
                  <router-view/>
                </Content>
              </Layout>
            </Content>
          </Layout>
          <!--<Footer class="layout-footer-center">2011-2016 &copy; TalkingData</Footer>-->
        </Layout>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      tabs: null
    }
  },
  created () {
    this.http.get(this.host + 'getCategory').then((response) => {
      this.tabs = response.data.data
    })
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
.ivu-layout-header {
  background-color: white;
  padding: 0;
}
  body {
    background-color: #f0f2f5;
  }
</style>
