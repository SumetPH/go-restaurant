<template>
  <div>
    <div class="row justify-content-center">
      <div class="col-md-6">
        <form @submit.prevent="find">
          <div class="form-group">
            <label>รหัสเมนู</label>
            <div class="row">
              <div class="col-md-10">
                <input v-model="MenuID" class="form-control" name="MenuID" type="text">
              </div>
              <div class="col-md-2">
                <button class="btn btn-primary" type="submit">ค้นหา</button>
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>
    <hr>
    <div class="row justify-content-center">
      <div class="col-md-6">
        <form @submit.prevent="edit">
          <div class="form-group">
            <label>รหัสเมนู</label>
            <input class="form-control" type="text" v-model="menu.MenuID" disabled>
          </div>
          <div class="form-group">
            <label>ชื่อเมนูอาหาร</label>
            <input class="form-control" type="text" v-model="menu.MenuName" required>
          </div>
          <div class="form-group">
            <label>ประเภทอาหาร</label>
            <select v-model="menu.MenuType" class="form-control" name="MenuType">
              <option value="1">อาหารคาว</option>
              <option value="2">อาหารหวาน</option>
              <option value="3">อาหารว่าง</option>
            </select>
          </div>
          <div class="form-group">
            <label>ราคา</label>
            <input class="form-control" type="number" v-model="menu.MenuPrice" required>
          </div>
          <div class="row justify-content-center">
            <button class="btn btn-warning" type="submit" :disabled="!menu">แก้ไขข้อมูล</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      MenuID: '',
      menu: false
    }
  },
  methods: {
    find() {
      let formData = new FormData()
      formData.append('id', this.MenuID)
      this.axios.post('/api/find', formData).then(res => {
        console.log(res.data)
        if (res.data.status === 'error') {
          this.menu = false
          alert(res.data.msg)
        }
        this.menu = res.data
      })
    },
    edit() {
      let formData = new FormData()
      formData.append('MenuID', this.menu.MenuID)
      formData.append('MenuName', this.menu.MenuName)
      formData.append('MenuType', this.menu.MenuType)
      formData.append('MenuPrice', this.menu.MenuPrice)

      this.axios.put('/api', formData).then(res => {
        console.log(res.data)
        if (res.data.status === 'error') {
          return alert(res.data.msg)
        }
        this.MenuID = ''
        this.menu = false
        return alert(res.data.msg)
      })
    }
  }
}
</script>
