<template>
  <div class="container">
    <slider></slider>
    <div id="term" style="font-family:  'Avenir', Helvetica, Arial, sans-serif;"></div>
  </div>
</template>
<script>
import $ from 'jquery'
import 'jquery.terminal/css/jquery.terminal.min.css'
import 'jquery.terminal/js/jquery.terminal.min.js'
export default {
  data () {
    return {
      http: this.http,
      host: this.host
    }
  },
  created () {
    $(function ($, http, host) {
      $('#term').terminal(function (command, term) {
        if (command !== '') {
          $.post('http://localhost:9092/' + 'outPut', {command: command}).then(function (response) {
            term.echo(response.data)
          })
        } else {
          this.echo('')
        }
      }, {
        greetings: 'Welcome To Tools, Just For Fun...\r\n@copyright toolsWebsite',
        name: 'ssh',
        height: 600,
        prompt: 'Tools@ForFun> '
      })
    })
  },
  methods: {
    exec (command) {
      this.http.post(this.host + 'outPut', {command: command}).then((response) => {
      })
    }
  }
}
</script>

<style scoped>
  .terminal span{
    font-family: 'Avenir', Helvetica, Arial, sans-serif !important;
  }
</style>
