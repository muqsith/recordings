import fetch from 'whatwg-fetch';

class DAL {
  constructor() {
    console.log('CONFIG: ', CONFIG);
  }

  async getRecordsList() {
    const result = await fetch();
    return result;
  }

  getRecord() {

  }

  createRecord() {

  }

  updateRecord() {
    
  }
 
}

export default DAL;