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
        <Upload
          ref="upload"
          :show-upload-list="false"
          :default-file-list="defaultList"
          :on-success="_handleSuccess"
          :format="['jpg','jpeg','png', 'gif']"
          :max-size="2048"
          :on-format-error="_handleFormatError"
          :on-exceeded-size="_handleMaxSize"
          :before-upload="handleBeforeUpload"
          multiple
          type="drag"
          :action="action"
          style="display: inline-block;width:58px;">
          <div style="width: 58px;height:58px;line-height: 58px;">
            <Icon type="ios-cloud-upload" size="20"></Icon>
          </div>
        </Upload>
        <Modal title="View Image" v-model="visible">
          <img :src="'https://o5wwk8baw.qnssl.com/' + imgName + '/large'" v-if="visible" style="width: 100%">
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
      action: this.host + 'upload'
    }
  },
  methods: {
    handleSuccess (res, file) {
      file.url = 'https://o5wwk8baw.qnssl.com/7eb99afb9d5f317c912f08b5212fd69a/avatar'
      file.name = '7eb99afb9d5f317c912f08b5212fd69a'
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
    }
  }
}
</script>

<style scoped>

</style>
