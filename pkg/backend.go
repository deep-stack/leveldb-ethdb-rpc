// VulcanizeDB
// Copyright © 2022 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package leveldb_ethdb_rpc

import (
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
)

var _ ethdb.Database = &LevelDBBackend{}

// NewLevelDBBackend creates a new levelDB RPC server backend
func NewLevelDBBackend(conf *Config) (*LevelDBBackend, error) {
	db, err := leveldb.New(conf.FilePath, conf.Cache, conf.Handles, conf.Namespace, true)
	if err != nil {
		return nil, err
	}
	frdb, err := rawdb.NewDatabaseWithFreezer(db, conf.FreezerPath, conf.Namespace, true)
	if err != nil {
		db.Close()
		return nil, err
	}
	return &LevelDBBackend{
		ethDB:   frdb,
		levelDB: db,
	}, nil
}

type LevelDBBackend struct {
	ethDB   ethdb.Database
	levelDB *leveldb.Database
}

func (s *LevelDBBackend) Has(key []byte) (bool, error) {
	return s.ethDB.Has(key)
}

func (s *LevelDBBackend) Get(key []byte) ([]byte, error) {
	return s.ethDB.Get(key)
}

func (s *LevelDBBackend) HasAncient(kind string, number uint64) (bool, error) {
	return s.ethDB.HasAncient(kind, number)
}

func (s *LevelDBBackend) Ancient(kind string, number uint64) ([]byte, error) {
	return s.ethDB.Ancient(kind, number)
}

func (s *LevelDBBackend) AncientRange(kind string, start, count, maxBytes uint64) ([][]byte, error) {
	return s.ethDB.AncientRange(kind, start, count, maxBytes)
}

func (s *LevelDBBackend) ReadAncients(fn func(ethdb.AncientReader) error) error {
	return s.ethDB.ReadAncients(fn)
}

func (s *LevelDBBackend) Ancients() (uint64, error) {
	return s.ethDB.Ancients()
}

func (s *LevelDBBackend) Tail() (uint64, error) {
	return s.ethDB.Tail()
}

func (s *LevelDBBackend) AncientSize(kind string) (uint64, error) {
	return s.ethDB.AncientSize(kind)
}

func (s *LevelDBBackend) Put(_ []byte, _ []byte) error {
	return errWriteNotAllowed
}

func (s *LevelDBBackend) Delete(_ []byte) error {
	return errWriteNotAllowed
}

func (s *LevelDBBackend) ModifyAncients(f func(ethdb.AncientWriteOp) error) (int64, error) {
	return 0, errWriteNotAllowed
}

func (s *LevelDBBackend) TruncateAncients(n uint64) error {
	return errWriteNotAllowed
}

func (s *LevelDBBackend) TruncateHead(n uint64) error {
	return errWriteNotAllowed
}

func (s *LevelDBBackend) TruncateTail(n uint64) error {
	return errWriteNotAllowed
}

func (s *LevelDBBackend) Sync() error {
	return errWriteNotAllowed
}

func (s *LevelDBBackend) NewBatch() ethdb.Batch {
	return nil
}

func (d *LevelDBBackend) NewBatchWithSize(size int) ethdb.Batch {
	return nil
}

func (s *LevelDBBackend) NewIterator(_ []byte, _ []byte) ethdb.Iterator {
	return nil
}

func (s *LevelDBBackend) Stat(property string) (string, error) {
	return s.ethDB.Stat(property)
}

func (s *LevelDBBackend) Compact(_ []byte, _ []byte) error {
	return errWriteNotAllowed
}

func (s *LevelDBBackend) Close() error {
	return errWriteNotAllowed
}

func (s *LevelDBBackend) MigrateTable(string, func([]byte) ([]byte, error)) error {
	return errWriteNotAllowed
}

func (s *LevelDBBackend) NewSnapshot() (ethdb.Snapshot, error) {
	return s.ethDB.NewSnapshot()
}
