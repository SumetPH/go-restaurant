<template>
  <div class="container">
    <div class="row m-5">
      <div class="col-md-12">
        <div class="card">
          <div class="card-header">
            <div class="row">
              <div class="col-md-10">
                <h3>หน้าหลักโปรแกรม</h3>
              </div>
              <div class="col-md-2 d-flex justify-content-center align-items-center">
                <router-link to="/">หน้าหลัก</router-link>
              </div>
            </div>
          </div>
          <div class="card-body">
            <div class="row justify-content-center m-3">
              <div class="col-md-2">
                <router-link to="/store">เพิ่มข้อมูล</router-link>
              </div>
              <div class="col-md-2">
                <router-link to="/search">ค้นหาข้อมูล</router-link>
              </div>
              <div class="col-md-2">
                <router-link to="/edit">แก้ไขข้อมูล</router-link>
              </div>
              <div class="col-md-2">
                <router-link to="/delete">ลบข้อมูล</router-link>
              </div>
            </div>
            <router-view/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
