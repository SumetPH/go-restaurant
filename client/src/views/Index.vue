<template>
  <div class="row">
    <div class="col-md-12">
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
</template>

<script>
export default {
  data() {
    return {
      menu: []
    }
  },
  mounted() {
    this.axios.get('/api').then(res => {
      console.log(res.data)
      this.menu = res.data.menu
    })
  }
}
</script>
