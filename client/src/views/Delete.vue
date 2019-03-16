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
        <form @submit.prevent="del">
          <div class="form-group">
            <label>รหัสเมนู</label>
            <input class="form-control" type="text" disabled v-model="menu.MenuID">
          </div>
          <div class="form-group">
            <label>ชื่อเมนูอาหาร</label>
            <input class="form-control" type="text" disabled v-model="menu.MenuName">
          </div>
          <div class="form-group">
            <label>ประเภทอาหาร</label>
            <input
              v-if="menu.MenuType === 1"
              class="form-control"
              type="text"
              disabled
              value="อาหารคาว"
            >
            <input
              v-else-if="menu.MenuType === 2"
              class="form-control"
              type="text"
              disabled
              value="อาหารหวาน"
            >
            <input
              v-else-if="menu.MenuType === 3"
              class="form-control"
              type="text"
              disabled
              value="อาหารว่าง"
            >
            <input v-else class="form-control" type="text" disabled>
          </div>
          <div class="form-group">
            <label>ราคา</label>
            <input class="form-control" type="number" disabled v-model="menu.MenuPrice">
          </div>
          <div class="row justify-content-center">
            <button class="btn btn-danger" type="submit" :disabled="!menu">ลบข้อมูล</button>
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
    del() {
      this.axios.delete(`/api/${this.menu.MenuID}`).then(res => {
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
