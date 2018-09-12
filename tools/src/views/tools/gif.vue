<template>
  <div class="container">
    <Tabs>
      <TabPane label="GIF编辑">
        <Upload
          ref="upload"
          :show-upload-list="false"
          :on-success="handleSuccess"
          :format="['gif']"
          :max-size="10240"
          :on-format-error="handleFormatError"
          :on-exceeded-size="handleMaxSize"
          type="drag"
          :action="action">
          <div style="padding: 20px 0">
            <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
            <p>点击上传gif图片,单张图片不超过10M</p>
          </div>
        </Upload>
        <Modal title="GIF 说明" v-model="visible" @on-ok="ok">
          <img :src="'http://img.com/' + imgName" v-if="visible" style="width: 100%">
          <Input placeholder="输入说明..." v-model="value" style="width: 100%;margin-top:15px;" />
        </Modal>
        <div class="demo-upload-list" :key="item.url" v-for="(item,index) in gifList" >
            <img :src="item.url">
            <div class="demo-upload-list-cover">
              <Icon type="ios-eye-outline" @click.native="handleView(item.name, index)"></Icon>
            </div>
        </div>
        <div class="container"  v-on:click="_produce()">
          <Button type="primary" >制作</Button>
        </div>
      </TabPane>
      <TabPane label="GIF合成">
        <div class="demo-upload-list" :key="item.url" v-for="item in uploadList" >
          <template v-if="item.status === 'finished'">
            <img :src="item.url">
            <div class="demo-upload-list-cover">
              <Icon type="ios-eye-outline" @click.native="handleView(item.name)"></Icon>
              <Icon type="ios-trash-outline" @click.native="handleRemove(item)"></Icon>
            </div>
          </template>
          <template v-else>
            <Progress v-if="item.showProgress" :percent="item.percentage" hide-info></Progress>
          </template>
        </div>
        <Upload
          ref="upload"
          :show-upload-list="false"
          :on-success="_handleSuccess"
          :format="['jpg','jpeg','png']"
          :max-size="2048"
          :on-format-error="_handleFormatError"
          :on-exceeded-size="_handleMaxSize"
          :before-upload="_handleBeforeUpload"
          multiple
          type="drag"
          :action="action"
          style="display: inline-block;width:58px;">
          <div style="width: 58px;height:58px;line-height: 58px;">
            <Icon type="ios-camera" size="20"></Icon>
          </div>
        </Upload>
        <div class="container"  v-on:click="produce()">
            <Button type="primary" @click="handleSpinCustom">制作</Button>
        </div>
        <Modal
          v-model="modal1"
          title="GIF">
          <p style="margin:0 auto;"><img :src="imgUrl"></p>
        </Modal>
      </TabPane>
    </Tabs>
  </div>
</template>

<script>
export default {
  name: 'gif',
  data () {
    return {
      action: this.host + 'upload',
      visible: false,
      imgName: '',
      uploadList: [],
      list: [],
      imgUrl: '',
      modal1: false,
      gifList: [],
      info: [],
      value: '',
      index: null
    }
  },
  methods: {
    handleSpinCustom () {
      this.$Spin.show({
        render: (h) => {
          return h('div', [
            h('Icon', {
              'class': 'demo-spin-icon-load',
              props: {
                type: 'ios-loading',
                size: 18
              }
            }),
            h('div', 'GIF Making Now...')
          ])
        }
      })
    },
    ok () {
      let item = []
      item.msg = this.value
      item.key = this.index
      this.info.push(item)
      this.value = ''
    },
    _produce () {
      let gif = ''
      let info = ''
      for (let item of this.gifList) {
        gif += item.name + ','
      }
      for (let item of this.info) {
        info += item.msg + '_' + item.key + ','
      }
      this.http.post(this.host + 'makeGifIntro', {gif: gif, info: info}).then((response) => {
      })
    },
    produce () {
      this.list = []
      for (let item of this.uploadList) {
        this.list.push(item.name)
      }
      this.http.post(this.host + 'makeGif', {list: this.list}).then((response) => {
        this.$Spin.hide()
        if (response.data.msg === 'success') {
          this.imgUrl = this.imgHost + response.data.data
          this.modal1 = true
          this.$Notice.success({
            title: 'GIF 合成',
            desc: 'GIF 合成成功'
          })
        }
      }).catch((error) => {
        this.$Spin.hide()
        this.$Notice.success({
          title: 'GIF 合成',
          desc: 'GIF 合成失败, 请重新执行'
        })
      })
    },
    handleSuccess (res, file) {
      this.gifList = []
      for (let v of res.data) {
        let item = []
        item.name = v
        item.url  = this.imgHost + v
        this.gifList.push(item)
      }
    },
    handleMaxSize (file) {
      this.$Notice.warning({
        title: '超过最大限制',
        desc: '文件最大支持 10M'
      })
    },
    handleFormatError (file) {
      this.$Notice.warning({
        title: '文件格式',
        desc: '只支持 GIF 图片格式'
      })
    },
    _handleFormatError (file) {
      this.$Notice.warning({
        title: '文件格式',
        desc: file.name + ' 格式错误, 只支持 JPG, JPEG, PNG, GIF 图片格式'
      })
    },
    _handleMaxSize (file) {
      this.$Notice.warning({
        title: '超过最大限制',
        desc: file.name + ' 文件过大, 最大支持 2M'
      })
    },
    _handleBeforeUpload () {
      const check = this.uploadList.length < 100
      if (!check) {
        this.$Notice.warning({
          title: '最多支持上传20张图片'
        })
      }
      return check
    },
    _handleSuccess (res, file) {
      file.url = this.imgHost + res.data
      file.name = res.data
    },
    handleView (name, index) {
      this.imgName = name
      this.visible = true
      this.index = index
    },
    handleRemove (file) {
      const fileList = this.$refs.upload.fileList
      this.$refs.upload.fileList.splice(fileList.indexOf(file), 1)
    }
  },
  mounted () {
    this.uploadList = this.$refs.upload.fileList
  }
}
</script>

<style>
  .demo-upload-list{
    display: inline-block;
    width: 60px;
    height: 60px;
    text-align: center;
    line-height: 60px;
    border: 1px solid transparent;
    border-radius: 4px;
    overflow: hidden;
    background: #fff;
    position: relative;
    box-shadow: 0 1px 1px rgba(0,0,0,.2);
    margin-right: 4px;
  }
  .demo-upload-list img{
    width: 100%;
    height: 100%;
  }
  .demo-upload-list-cover{
    display: none;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: rgba(0,0,0,.6);
  }
  .demo-upload-list:hover .demo-upload-list-cover{
    display: block;
  }
  .demo-upload-list-cover i{
    color: #fff;
    font-size: 20px;
    cursor: pointer;
    margin: 0 2px;
  }
  .demo-spin-icon-load{
    animation: ani-demo-spin 1s linear infinite;
  }
</style>
