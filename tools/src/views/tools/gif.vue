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
      </TabPane>
      <TabPane label="GIF合成">
        <div class="demo-upload-list" v-for="item in uploadList" >
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
        <Modal title="View Image" v-model="visible">
          <img :src="'http://img.com/' + imgName" v-if="visible" style="width: 100%">
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
      uploadList: []
    }
  },
  methods: {
    handleSuccess (res, file) {
      console.log(123)
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
      const check = this.uploadList.length < 50
      if (!check) {
        this.$Notice.warning({
          title: '最多支持上传20张图片'
        })
      }
      return check
    },
    _handleSuccess (res, file) {
      file.url = 'http://img.com/' + res.data
      file.name = res.data
    },
    handleView (name) {
      this.imgName = name
      this.visible = true
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
</style>
