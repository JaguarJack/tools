<template>
  <div class="container" style="margin-top:10px;">

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
    this.load()
  },
  watch: {
    '$route': 'load'
  },
  methods: {
    load () {
      this.http.post(this.host + 'getTools', {name: this.$route.params.code}).then((response) => {
        this.tools = response.data.data
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
