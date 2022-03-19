import * as React from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";

import { getAlbumsList } from "../../DAL/index";

export const List = () => {
  const [albums, setAlbums] = React.useState([]);

  const fetchAlbums = async () => {
    const albumsList = await getAlbumsList();
    setAlbums(albumsList);
  };

  React.useEffect(() => {
    fetchAlbums();
  }, []);

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>ID</TableCell>
            <TableCell>Title</TableCell>
            <TableCell>Artist</TableCell>
            <TableCell align="right">Price USD</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {albums.map((album) => (
            <TableRow
              key={album.id}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell component="th" scope="row">
                {album.id}
              </TableCell>
              <TableCell component="th" scope="row">
                {album.title}
              </TableCell>
              <TableCell component="th" scope="row">
                {album.artist}
              </TableCell>
              <TableCell align="right">{album.price}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};
