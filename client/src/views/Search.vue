<template>
  <div>
    <div class="row justify-content-center">
      <div class="col-md-6">
        <form @submit.prevent="search">
          <div class="form-group">
            <label>รหัส/ชื่อเมนูอาหาร</label>
            <div class="row">
              <div class="col-md-10">
                <input v-model="key" class="form-control" name="key" type="text">
              </div>
              <div class="col-md-2">
                <button class="btn btn-primary" type="submit">ค้นหา</button>
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>
    <div class="row justify-content-center mt-3">
      <div class="col-md-10">
        <table class="table">
          <thead>
            <tr>
              <th>รหัสเมนู</th>
              <th>ชื่อเมนูอาหาร</th>
              <th>ประเภทอาหาร</th>
              <th>ราคา</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in menu" :key="index">
              <td>{{ item.MenuID }}</td>
              <td>{{ item.MenuName }}</td>
              <td v-if="item.MenuType === 1">อาหารคาว</td>
              <td v-if="item.MenuType === 2">อาหารหวาน</td>
              <td v-if="item.MenuType === 3">อาหารว่าง</td>
              <td>{{ item.MenuPrice }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      key: '',
      menu: []
    }
  },
  methods: {
    search() {
      let formData = new FormData()
      formData.append('key', this.key)
      this.axios.post('/api/search', formData).then(res => {
        console.log(res.data)
        this.menu = res.data.menu
      })
    }
  }
}
</script>