<template>
  <div class="row justify-content-center">
    <div class="col-md-9">
      <form @submit.prevent="submit">
        <div class="form-group">
          <label>รหัสเมนู</label>
          <input
            v-model="MenuID"
            class="form-control"
            name="MenuID"
            type="text"
            maxlength="5"
            required
          >
        </div>
        <div class="form-group">
          <label>ชื่อเมนูอาหาร</label>
          <input
            v-model="MenuName"
            class="form-control"
            name="MenuName"
            type="text"
            maxlength="50"
            required
          >
        </div>
        <div class="form-group">
          <label>ประเภทอาหาร</label>
          <select v-model="MenuType" class="form-control" name="MenuType">
            <option value="1">อาหารคาว</option>
            <option value="2">อาหารหวาน</option>
            <option value="3">อาหารว่าง</option>
          </select>
        </div>
        <div class="form-group">
          <label>ราคา</label>
          <input v-model="MenuPrice" class="form-control" name="MenuPrice" type="number" max="9999">
        </div>
        <div class="row justify-content-center">
          <button class="btn btn-primary" type="submit">เพิ่มข้อมูล</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      MenuID: '',
      MenuName: '',
      MenuType: 1,
      MenuPrice: ''
    }
  },
  methods: {
    submit() {
      let formData = new FormData()
      formData.append('MenuID', this.MenuID)
      formData.append('MenuName', this.MenuName)
      formData.append('MenuType', this.MenuType)
      formData.append('MenuPrice', this.MenuPrice)
      this.axios.post('/api', formData).then(res => {
        console.log(res.data)
        if (res.data.status === 'success') {
          this.MenuID = ''
          this.MenuName = ''
          this.MenuType = 1
          this.MenuPrice = ''
          return alert(res.data.msg)
        } else {
          return alert(res.data.msg)
        }
      })
    }
  }
}
</script>
