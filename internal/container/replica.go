package container

import (
	"fmt"

	"github.com/AndyS1mpson/key-value-database/internal/database/storage/replication"
	"github.com/AndyS1mpson/key-value-database/internal/infrastructure/network/tcp_client"
	"github.com/AndyS1mpson/key-value-database/internal/infrastructure/network/tcp_server"
	"github.com/AndyS1mpson/key-value-database/internal/utils/container"
	utils "github.com/AndyS1mpson/key-value-database/internal/utils/size_parser"
)

func (c *Container) getMasterReplica() *replication.Master {
	return container.MustOrGetNew(c.Container, func() *replication.Master {
		var options []tcp_server.Option
		options = append(options, tcp_server.WithServerIdleTimeout(c.config.Replication.SyncInterval*3))
		options = append(options, tcp_server.WithServerMaxConnectionsNumber(uint(c.config.Replication.MaxReplicasNumber)))

		if c.config.WAL.MaxSegmentSize != "" {
			bufferSize, err := utils.ParseSize(c.config.WAL.MaxSegmentSize)
			if err != nil {
				panic(fmt.Sprintf("can not parse WAL max segment size: %s", err))
			}

			options = append(options, tcp_server.WithServerBufferSize(uint(bufferSize)))
		}

		server, err := tcp_server.NewTCPServer(c.config.Replication.MasterAddress, c.GetLogger(), options...)
		if err != nil {
			panic(fmt.Sprintf("failed to create tcp server: %s", err))
		}

		replica := replication.NewMaster(server, c.config.WAL.DirectoryPath, c.GetLogger())

		go replica.Start(c.Ctx())

		return replica
	})
}

func (c *Container) getSlaveReplica() *replication.Slave {
	return container.MustOrGetNew(c.Container, func() *replication.Slave {
		var options []tcp_client.Option
		options = append(options, tcp_client.WithClientIdleTimeout(c.config.Replication.SyncInterval*3))

		if c.config.WAL.MaxSegmentSize != "" {
			bufferSize, err := utils.ParseSize(c.config.WAL.MaxSegmentSize)
			if err != nil {
				panic(fmt.Sprintf("can not parse WAL max segment size: %s", err))
			}

			options = append(options, tcp_client.WithClientBufferSize(uint(bufferSize)))
		}

		client, err := tcp_client.NewTCPClient(c.config.Replication.MasterAddress, options...)
		if err != nil {
			panic(fmt.Sprintf("failed to create tcp client: %s", err))
		}

		replica := replication.NewSlave(
			client,
			c.config.WAL.DirectoryPath,
			c.config.Replication.SyncInterval,
			c.GetLogger(),
		)

		go replica.Start(c.Ctx())

		return replica
	})
}
