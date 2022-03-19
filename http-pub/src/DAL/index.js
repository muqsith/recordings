import "whatwg-fetch";

// CONFIG is global and it is injected by webpack with DefinePlugin
const apiUrl = CONFIG.apiUrl;

export const getAlbumsList = async () => {
  const response = await fetch(apiUrl + "/api/albums");
  const result = await response.json();
  return result;
};

export const getAlbum = async () => {};

export const createAlbum = async () => {};

export const updateAlbum = async () => {};
